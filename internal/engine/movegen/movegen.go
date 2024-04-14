package movegen

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

// TODO: Удаляется вертикаль??? Исправить
// TODO: GameState сделать указателем

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

func generateGameStateAfterChangeAt(board board.Board, nextPlayer model.Player, changePos board.Position) *GameState {
	ipUnderAttack := board.IsKingAttackedByPiece(changePos)

	return &GameState{Board: board, CurMovePlayer: nextPlayer, KingUnderAttack: ipUnderAttack}
}

func tryGenerateGameStateAfterChangeAtWithNextMovesCheck(board board.Board, nextPlayer model.Player, changePos board.Position) *GameState {
	// e.g., pawn cannot move to last row because it will have no moves to board field (so this is not possible game state)
	if len(board.GetPieceMovesToBoardField(changePos)) != 0 {
		return generateGameStateAfterChangeAt(board, nextPlayer, changePos)
	}
	return nil
}

func tryGeneratePromotionGameStateWithMove(someBoard board.Board, nextPlayer model.Player, fromPos, toPos board.Position) *GameState {
	fromFile, fromRank := fromPos.Get()
	toFile, toRank := toPos.Get()
	var boardPiece = someBoard.Cells[fromFile][fromRank]
	if boardPiece.IsPromotable() {
		var toPromotionZone = slices.Contains(someBoard.GetPromotionZone(boardPiece.Player), toRank)
		var fromPromotionZone = slices.Contains(someBoard.GetPromotionZone(boardPiece.Player), fromRank)
		if toPromotionZone || fromPromotionZone {
			altNewBoard := someBoard.Clone()
			altNewBoard.Cells[toFile][toRank] = boardPiece.GetPromotedPiece()
			altNewBoard.Cells[fromFile][fromRank] = nil

			return generateGameStateAfterChangeAt(altNewBoard, nextPlayer, toPos)
		}
	}
	return nil
}

func (gs *GameState) generatePossibleStatesWithMove(fromPos, toPos board.Position) (gss []GameState) {
	gss = []GameState{}
	fromFile, fromRank := fromPos.Get()
	toFile, toRank := toPos.Get()

	var newBoard = gs.Board.Clone()
	var nextPlayer = gs.getNextPlayer()

	var cellMoveTo = newBoard.Cells[toFile][toRank]
	var emptyCell = cellMoveTo == nil
	if !emptyCell {
		newBoard.Inventories[gs.CurMovePlayer].AddPiece(cellMoveTo)
	}

	// creating alternative gamestate for promotion of piece
	promotionGameState := tryGeneratePromotionGameStateWithMove(newBoard, nextPlayer, fromPos, toPos)
	if promotionGameState != nil {
		gss = append(gss, *promotionGameState)
	}

	newBoard.Cells[toFile][toRank] = newBoard.Cells[fromFile][fromRank]
	newBoard.Cells[fromFile][fromRank] = nil

	newGameState := tryGenerateGameStateAfterChangeAtWithNextMovesCheck(newBoard, nextPlayer, toPos)
	if newGameState != nil {
		gss = append(gss, *newGameState)
	}

	return
}

func (gs *GameState) generatePossibleStatesFromBoardPiece(piecePos board.Position) (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var movesPositions = curBoard.GetPiecePossibleMoves(piecePos)
	for _, movePos := range movesPositions {
		gss = append(gss, gs.generatePossibleStatesWithMove(piecePos, movePos)...)
	}
	return
}

func (gs *GameState) generatePossibleStatesFromBoardPieces() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	curBoard.IterateBoardPieces(gs.CurMovePlayer, func(piece *board.Piece, pos board.Position) {
		gss = append(gss, gs.generatePossibleStatesFromBoardPiece(pos)...)
	})
	return
}

func (gs *GameState) generatePossibleStatesFromKingRunningAway(kingPos board.Position) (gss []GameState) {
	gss = []GameState{}
	var curBoard = gs.Board
	var movesPositions = curBoard.GetKingMoves(kingPos)
	for _, movePos := range movesPositions {
		gss = append(gss, gs.generatePossibleStatesWithMove(kingPos, movePos)...)
	}
	return
}

func (gs *GameState) generatePossibleStatesWithDrop(pieceType *model.PieceType, dropPos board.Position) (gss []GameState) {
	gss = []GameState{}
	dropFile, dropRank := dropPos.Get()

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer
	var nextPlayer = gs.getNextPlayer()

	// check for two pawns in a column
	if pieceType.Name == "Pawn" && curBoard.IsTherePawn(dropFile) {
		return
	}

	newBoard := curBoard.Clone()
	newBoard.Cells[dropFile][dropRank] = newBoard.Inventories[curPlayer].ExtractPieceToPlayer(pieceType, curPlayer)

	if pieceType.Name == "Pawn" {
		pawnMovesPositions := newBoard.GetPiecePossibleMoves(dropPos)
		for _, movePos := range pawnMovesPositions {
			moveFile, moveRank := movePos.Get()
			var cell = newBoard.Cells[moveFile][moveRank]
			if cell != nil && cell.Type.ImportantPiece {
				var runningFromPawnStates = gs.generatePossibleStatesFromKingRunningAway(movePos)
				if len(runningFromPawnStates) == 0 { // if king can't run away from dropped pawn
					return
				}
			}
		}
	}

	newGameState := tryGenerateGameStateAfterChangeAtWithNextMovesCheck(newBoard, nextPlayer, dropPos)
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

	var dropsPositions = curBoard.GetPossibleDropsCoords()
	for _, dropPos := range dropsPositions {
		for _, piece := range curBoard.Inventories[gs.CurMovePlayer].Pieces() {
			gss = append(gss, gs.generatePossibleStatesWithDrop(piece, dropPos)...)
		}
	}
	return
}

func (gs *GameState) generatePossibleStatesFromDefendingKing() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer

	var kingPosition = curBoard.GetKingPositionForPlayer(curPlayer)
	var kingFile, kingRank = kingPosition.Get()
	var king = curBoard.Cells[kingFile][kingRank]

	var attackerPlayer = king.GetAttackerPlayer()
	var attackers = curBoard.GetPositionsOfAttackersOnCell(attackerPlayer, kingPosition)

	if len(attackers) == 1 { // can eat attacker or drop(move) piece on his path
		attackerPos := attackers[0]
		attackerPoint := graphics.NewPoint(attackerPos.GetFile(), attackerPos.GetRank())
		ipPoint := graphics.NewPoint(kingFile, kingRank)
		attackPath := graphics.GetLinePoints(attackerPoint, ipPoint) // TODO: будет неправильно работать для коня, исправить

		// eating and closing from
		curBoard.IterateBoardPieces(curPlayer, func(piece *board.Piece, pos board.Position) {
			// this condition need to not creating two same game states
			// (because important piece will eat attacker when it starts to run away (see below))
			// + to not create game state where important piece closing by itself
			if !piece.Type.ImportantPiece {
				var movesPositions = curBoard.GetPiecePossibleMoves(pos)

				// eating
				idx := slices.Index(movesPositions, attackerPos)
				if idx != -1 {
					gss = append(gss, gs.generatePossibleStatesWithMove(pos, attackerPos)...)
				}

				// closing from
				for i := range movesPositions {
					moveFile, moveRank := movesPositions[i].Get()
					for j := 1; j < len(attackPath)-1; j++ {
						closeFile, closeRank := attackPath[j].Coordinates()
						if closeFile == moveFile && closeRank == moveRank {
							gss = append(gss, gs.generatePossibleStatesWithMove(pos, board.NewPos(closeFile, closeRank))...)
							break
						}
					}
				}
			}
		})

		// droping
		if !curBoard.Inventories[curPlayer].IsEmpty() {
			for i := 1; i < len(attackPath)-1; i++ {
				for _, pieceType := range curBoard.Inventories[curPlayer].Pieces() {
					dropFile, dropRank := attackPath[i].Coordinates()
					gss = append(gss, gs.generatePossibleStatesWithDrop(pieceType, board.NewPos(dropFile, dropRank))...)
				}
			}
		}
	}

	// important piece is running away
	gss = append(gss, gs.generatePossibleStatesFromKingRunningAway(kingPosition)...)
	return
}

func (gs *GameState) GeneratePossibleStates() (gss []GameState) {
	gss = []GameState{}

	if !gs.KingUnderAttack {
		gss = append(gss, gs.generatePossibleStatesFromBoardPieces()...)
		gss = append(gss, gs.generatePossibleStatesFromInventoryPieces()...)
	} else {
		gss = append(gss, gs.generatePossibleStatesFromDefendingKing()...)
	}

	return
}
