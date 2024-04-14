package movegen

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

// TODO: Везде поменять vertical и horizontal на file и rank (и добавить структуру Position (и мб структуру Move))
// TODO: Короля могут съесть, если он сделает ход на опасную клетку. Исправить
// TODO: Удаляется вертикаль??? Исправить

type GameState struct {
	Board           board.Board
	CurMovePlayer   model.Player
	KingUnderAttack bool
}

func (gs *GameState) getNextPlayer() model.Player {
	if gs.CurMovePlayer == model.Sente {
		return model.Gote
	} else {
		return model.Sente
	}
}

func generateGameStateAfterChangeAt(board board.Board, nextPlayer model.Player, vChangeCoord int, hChangeCoord int) *GameState {
	ipUnderAttack := board.IsKingAttackedByPiece(vChangeCoord, hChangeCoord)

	return &GameState{Board: board, CurMovePlayer: nextPlayer, KingUnderAttack: ipUnderAttack}
}

func tryGenerateGameStateAfterChangeAtWithNextMovesCheck(board board.Board, nextPlayer model.Player, vChangeCoord int, hChangeCoord int) *GameState {
	// e.g., pawn cannot move to last row because it will have no moves to board field (so this is not possible game state)
	if len(board.GetInBoardFieldMoves(vChangeCoord, hChangeCoord)) != 0 {
		return generateGameStateAfterChangeAt(board, nextPlayer, vChangeCoord, hChangeCoord)
	}
	return nil
}

func tryGeneratePromotionGameStateWithMove(someBoard board.Board, nextPlayer model.Player, vCoordFrom int, hCoordFrom int, vCoordTo int, hCoordTo int) *GameState {
	var boardPiece = someBoard.Cells[vCoordFrom][hCoordFrom]
	if boardPiece.IsPromotable() {
		var toPromotionZone = slices.Contains(someBoard.GetPromotionZone(boardPiece.Player), hCoordTo)
		var fromPromotionZone = slices.Contains(someBoard.GetPromotionZone(boardPiece.Player), hCoordFrom)
		if toPromotionZone || fromPromotionZone {
			altNewBoard := someBoard.Clone()
			altNewBoard.Cells[vCoordTo][hCoordTo] = &board.Piece{Type: *boardPiece.Type.PromotePiece, Player: boardPiece.Player}
			altNewBoard.Cells[vCoordFrom][hCoordFrom] = nil

			return generateGameStateAfterChangeAt(altNewBoard, nextPlayer, vCoordTo, hCoordTo)
		}
	}
	return nil
}

func (gs *GameState) generatePossibleStatesWithMove(vCoordFrom int, hCoordFrom int, vCoordTo int, hCoordTo int) (gss []GameState) {
	gss = []GameState{}

	var newBoard = gs.Board.Clone()
	var nextPlayer = gs.getNextPlayer()

	var cellMoveTo = newBoard.Cells[vCoordTo][hCoordTo]
	var emptyCell = cellMoveTo == nil
	if !emptyCell {
		newBoard.Inventories[gs.CurMovePlayer].AddPiece(cellMoveTo)
	}

	// creating alternative gamestate for promotion of piece
	promotionGameState := tryGeneratePromotionGameStateWithMove(newBoard, nextPlayer, vCoordFrom, hCoordFrom, vCoordTo, hCoordTo)
	if promotionGameState != nil {
		gss = append(gss, *promotionGameState)
	}

	newBoard.Cells[vCoordTo][hCoordTo] = newBoard.Cells[vCoordFrom][hCoordFrom]
	newBoard.Cells[vCoordFrom][hCoordFrom] = nil

	newGameState := tryGenerateGameStateAfterChangeAtWithNextMovesCheck(newBoard, nextPlayer, vCoordTo, hCoordTo)
	if newGameState != nil {
		gss = append(gss, *newGameState)
	}

	return
}

func (gs *GameState) generatePossibleStatesFromBoardPiece(verticalCoord int, horizontalCoord int) (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var movesCoords [][2]int = curBoard.GetPossibleMovesCoords(verticalCoord, horizontalCoord)
	for _, coords := range movesCoords {
		gss = append(gss, gs.generatePossibleStatesWithMove(verticalCoord, horizontalCoord, coords[0], coords[1])...)
	}
	return
}

func (gs *GameState) generatePossibleStatesFromBoardPieces() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	for v, verticalCells := range curBoard.Cells {
		for h, cell := range verticalCells {
			if cell == nil { // empty cell
				continue
			}

			if cell.Player == gs.CurMovePlayer {
				gss = append(gss, gs.generatePossibleStatesFromBoardPiece(v, h)...)
			}
		}
	}
	return
}

func (gs *GameState) generatePossibleStatesWithDrop(pieceType *model.PieceType, vCoordTo int, hCoordTo int) (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer
	var nextPlayer = gs.getNextPlayer()

	// check for two pawns in a column
	if pieceType.Name == "Pawn" && curBoard.IsTherePawn(vCoordTo) {
		return
	}

	newBoard := curBoard.Clone()
	newBoard.Cells[vCoordTo][hCoordTo] = newBoard.Inventories[curPlayer].ExtractPieceToPlayer(pieceType, curPlayer)

	if pieceType.Name == "Pawn" {
		pawnMovesCoords := newBoard.GetPossibleMovesCoords(vCoordTo, hCoordTo)
		for _, coords := range pawnMovesCoords {
			var cell = newBoard.Cells[coords[0]][coords[1]]
			if cell != nil && cell.Type.ImportantPiece {
				var runningFromPawnStates = gs.generatePossibleStatesFromKingRunningAway(coords[0], coords[1])
				if len(runningFromPawnStates) == 0 { // if king can't run away from dropped pawn
					return
				}
			}
		}
	}

	newGameState := tryGenerateGameStateAfterChangeAtWithNextMovesCheck(newBoard, nextPlayer, vCoordTo, hCoordTo)
	if newGameState != nil {
		gss = append(gss, *newGameState)
	}
	return
}

func (gs *GameState) generatePossibleStatesFromInventoryPieces() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer

	if curBoard.Inventories[curPlayer].IsEmpty() {
		return
	}

	var dropsCoords = curBoard.GetPossibleDropsCoords()
	for _, coords := range dropsCoords {
		for _, pieceType := range curBoard.Inventories[gs.CurMovePlayer].Pieces() {
			gss = append(gss, gs.generatePossibleStatesWithDrop(pieceType, coords[0], coords[1])...)
		}
	}
	return
}

func (gs *GameState) generatePossibleStatesFromKingRunningAway(vIPCoord int, hIPCoord int) (gss []GameState) {
	gss = []GameState{}
	var curBoard = gs.Board
	var potentialSaveCoords = curBoard.GetPossibleMovesCoords(vIPCoord, hIPCoord)
	for _, coordsTo := range potentialSaveCoords {
		attackerMoves := curBoard.GetMovesCoordsOfAttackersOnCell(coordsTo[0], coordsTo[1])
		var isDangerCoords bool = false
		for i := range attackerMoves {
			if attackerMoves[i] == coordsTo {
				isDangerCoords = true
				break
			}
		}

		if !isDangerCoords {
			gss = append(gss, gs.generatePossibleStatesWithMove(vIPCoord, hIPCoord, coordsTo[0], coordsTo[1])...)
		}
	}
	return
}

func (gs *GameState) generatePossibleStatesFromDefendingKing() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer
	var ipCoords = curBoard.GetKingCoordsForPlayer(curPlayer)

	var attackers = curBoard.GetCoordsOfAttackersOnCell(ipCoords[0], ipCoords[1])

	if len(attackers) == 1 { // can eat attacker or drop(move) piece on his path
		attacker := attackers[0]
		attackerPoint := graphics.NewPoint(attacker[0], attacker[1])
		ipPoint := graphics.NewPoint(ipCoords[0], ipCoords[1])
		attackPath := graphics.GetLinePoints(attackerPoint, ipPoint) // TODO: будет неправильно работать для коня, исправить

		// eating and closing from
		for v := range curBoard.Cells {
			for h, cell := range curBoard.Cells[v] {
				if cell == nil { // empty cell
					continue
				}

				// second condition need to not creating two same game states
				// (because important piece will eat attacker when it starts to run away (see below))
				if cell.Player == curPlayer && !cell.Type.ImportantPiece {
					var movesCoords = curBoard.GetPossibleMovesCoords(v, h)

					// eating
					idx := slices.Index(movesCoords, [2]int{attacker[0], attacker[1]})
					if idx != -1 {
						gss = append(gss, gs.generatePossibleStatesWithMove(v, h, attacker[0], attacker[1])...)
					}

					// closing from
					for i := range movesCoords {
						for j := 1; j < len(attackPath)-1; j++ {
							vCloseCoord, hCloseCoord := attackPath[j].Coordinates()
							if [2]int{vCloseCoord, hCloseCoord} == movesCoords[i] {
								gss = append(gss, gs.generatePossibleStatesWithMove(v, h, vCloseCoord, hCloseCoord)...)
								break
							}
						}
					}
				}
			}
		}

		// droping
		if !curBoard.Inventories[curPlayer].IsEmpty() {
			for i := 1; i < len(attackPath)-1; i++ {
				for _, pieceType := range curBoard.Inventories[curPlayer].Pieces() {
					vCoordDropTo, hCoordDropTo := attackPath[i].Coordinates()
					gss = append(gss, gs.generatePossibleStatesWithDrop(pieceType, vCoordDropTo, hCoordDropTo)...)
				}
			}
		}
	}

	// important piece is running away
	gss = append(gss, gs.generatePossibleStatesFromKingRunningAway(ipCoords[0], ipCoords[1])...)
	return
}

func (gs *GameState) GetPossibleStates() (gss []GameState) {
	gss = []GameState{}

	if !gs.KingUnderAttack {
		gss = append(gss, gs.generatePossibleStatesFromBoardPieces()...)
		gss = append(gss, gs.generatePossibleStatesFromInventoryPieces()...)
	} else {
		gss = append(gss, gs.generatePossibleStatesFromDefendingKing()...)
	}

	return
}
