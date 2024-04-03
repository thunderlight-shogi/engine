package movegen

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

//TODO: Везде поменять vertical и horizontal на file и rank (и добавить структуру Position)
//TODO: Обработать правило: нельзя ставить мат пешкой с руки

type GameState struct {
	Board                     board.Board
	CurMovePlayer             model.Player
	ImportantPieceUnderAttack bool
	NumOfAttackableCells      uint //TODO: Сделать подсчет
	NumOfDropableCells        uint //TODO: Сделать подсчет
}

func (gs *GameState) getNextPlayer() model.Player {
	if gs.CurMovePlayer == model.Sente {
		return model.Gote
	} else {
		return model.Sente
	}
}

func generateGameStateAfterChangeAt(board board.Board, nextPlayer model.Player, vChangeCoord int, hChangeCoord int) *GameState {
	ipUnderAttack := board.IsImportantPieceAttackedByPiece(vChangeCoord, hChangeCoord)

	return &GameState{Board: board, CurMovePlayer: nextPlayer, ImportantPieceUnderAttack: ipUnderAttack}
}

func tryGenerateGameStateAfterChangeAtWithNextMovesCheck(board board.Board, nextPlayer model.Player, vChangeCoord int, hChangeCoord int) *GameState {
	// e. g., pawn cannot move to last row because it will have no moves to board field (so this is not possible game state)
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

func (gs *GameState) generatePossibleStatesFromDefendingImportantPiece() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer
	var ipCoords = curBoard.GetImportantPieceCoordsForPlayer(curPlayer)

	var attackers = curBoard.GetCoordsOfAttackersOnCell(ipCoords[0], ipCoords[1])
	var attackersMovesCoords [][2]int
	for i := range attackers {
		attackersMovesCoords = append(attackersMovesCoords, curBoard.GetPossibleMovesCoords(attackers[i][0], attackers[i][1])...)
	}

	if len(attackers) == 1 { //can eat attacker or drop(move) piece on his path
		attacker := attackers[0]
		attackerPoint := graphics.NewPoint(attacker[0], attacker[1])
		ipPoint := graphics.NewPoint(ipCoords[0], ipCoords[1])
		attackPath := graphics.GetLinePoints(attackerPoint, ipPoint)

		//eating and closing from
		for v := range curBoard.Cells {
			for h, cell := range curBoard.Cells[v] {
				if cell == nil { // empty cell
					continue
				}

				//second condition need to not creating two same game states
				//(because important piece will eat attacker when it starts to run away (see below))
				if cell.Player == curPlayer && !cell.Type.ImportantPiece {
					var movesCoords = curBoard.GetPossibleMovesCoords(v, h)

					//eating
					idx := slices.Index(movesCoords, [2]int{attacker[0], attacker[1]})
					if idx != -1 {
						gss = append(gss, gs.generatePossibleStatesWithMove(v, h, attacker[0], attacker[1])...)
					}

					//closing from
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

		//droping
		if !curBoard.Inventories[curPlayer].IsEmpty() {
			for i := 1; i < len(attackPath)-1; i++ {
				for _, pieceType := range curBoard.Inventories[curPlayer].Pieces() {
					vCoordDropTo, hCoordDropTo := attackPath[i].Coordinates()
					gss = append(gss, gs.generatePossibleStatesWithDrop(pieceType, vCoordDropTo, hCoordDropTo)...)
				}
			}
		}
	}

	//important piece is running away
	var potentialSaveCoords = curBoard.GetPossibleMovesCoords(ipCoords[0], ipCoords[1])
	for _, coordsTo := range potentialSaveCoords {

		var isDangerCoords bool = false
		for i := range attackersMovesCoords {
			if attackersMovesCoords[i] == coordsTo {
				isDangerCoords = true
			}
		}

		if !isDangerCoords {
			gss = append(gss, gs.generatePossibleStatesWithMove(ipCoords[0], ipCoords[1], coordsTo[0], coordsTo[1])...)
		}
	}

	return
}

func (gs *GameState) GetPossibleStates() (gss []GameState) {
	gss = []GameState{}

	if !gs.ImportantPieceUnderAttack {
		gss = append(gss, gs.generatePossibleStatesFromBoardPieces()...)
		gss = append(gss, gs.generatePossibleStatesFromInventoryPieces()...)
	} else {
		gss = append(gss, gs.generatePossibleStatesFromDefendingImportantPiece()...)
	}

	return
}
