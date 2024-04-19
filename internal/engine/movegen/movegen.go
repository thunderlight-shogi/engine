package movegen

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/gamestate"
	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

// TODO: Удаляется вертикаль??? Исправить
// TODO: GameState сделать указателем
// TODO: Добавить на вход GeneratePossibleMoves переменную, которая будет определять
//       нужно ли кешировать возможные ходы фигур для генерируемых досок

// TODO: Убрать changePos и возможно перенести эту функцию в gamestate (типо конструктор)
func generateGameStateAfterChangeAt(board board.Board, nextPlayer model.Player, changePos board.Position) *gamestate.GameState {
	ipUnderAttack := board.IsKingAttacked(nextPlayer)

	return &gamestate.GameState{Board: board, CurMovePlayer: nextPlayer, KingUnderAttack: ipUnderAttack}
}

func tryGenerateGameStateAfterChangeAtWithNextMovesCheck(someBoard board.Board, curPlayer, nextPlayer model.Player, changePos board.Position) *gamestate.GameState {
	// e.g., pawn cannot move to last row because it will have no moves to board field (so this is not possible game state)
	if len(someBoard.GetPieceMovesToBoardField(changePos)) != 0 {
		// TODO: сильно замедляет, исправить
		var isCurPlayerKingAttacked = someBoard.IsKingAttacked(curPlayer)

		if !isCurPlayerKingAttacked {
			return generateGameStateAfterChangeAt(someBoard, nextPlayer, changePos)
		}
	}
	return nil
}

func tryGeneratePromotionGameStateWithMove(someBoard board.Board, curPlayer, nextPlayer model.Player, fromPos, toPos board.Position) *gamestate.GameState {
	fromRank := fromPos.Rank
	toRank := toPos.Rank
	var boardPiece = someBoard.At(fromPos)
	if boardPiece.IsPromotable() {
		var toPromotionZone = slices.Contains(someBoard.GetPromotionZone(boardPiece.Player), toRank)
		var fromPromotionZone = slices.Contains(someBoard.GetPromotionZone(boardPiece.Player), fromRank)
		if toPromotionZone || fromPromotionZone {
			altNewBoard := someBoard.Clone()
			altNewBoard.MakeMove(fromPos, toPos, true)

			// TODO: сильно замедляет, исправить
			var isCurPlayerKingAttacked = someBoard.IsKingAttacked(curPlayer)

			if !isCurPlayerKingAttacked {
				return generateGameStateAfterChangeAt(altNewBoard, nextPlayer, toPos)
			}
		}
	}
	return nil
}

func generatePossibleStatesWithMove(gs *gamestate.GameState, fromPos, toPos board.Position) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	var newBoard = gs.Board.Clone()
	var nextPlayer = gs.GetNextPlayer()

	// creating alternative gamestate for promotion of piece
	promotionGameState := tryGeneratePromotionGameStateWithMove(newBoard, gs.CurMovePlayer, nextPlayer, fromPos, toPos)
	if promotionGameState != nil {
		gss = append(gss, *promotionGameState)
	}

	newBoard.MakeMove(fromPos, toPos, false)

	newGameState := tryGenerateGameStateAfterChangeAtWithNextMovesCheck(newBoard, gs.CurMovePlayer, nextPlayer, toPos)
	if newGameState != nil {
		gss = append(gss, *newGameState)
	}
	return
}

func generatePossibleStatesWithDrop(gs *gamestate.GameState, pieceType *model.PieceType, dropPos board.Position) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer
	var nextPlayer = gs.GetNextPlayer()

	// check for two pawns in a column
	if pieceType.Name == "Pawn" && curBoard.IsTherePawn(dropPos.File) {
		return
	}

	newBoard := curBoard.Clone()
	newBoard.MakeDrop(pieceType, curPlayer, dropPos)

	// TODO: по-любому как-то можно оптимизировать сейчас
	if pieceType.Name == "Pawn" {
		pawnMovesPositions := newBoard.GetPiecePossibleMoves(dropPos, true)
		for _, movePos := range pawnMovesPositions {
			var cell = newBoard.At(movePos)
			if cell != nil && cell.Type.ImportantPiece {
				var runningFromPawnStates = generatePossibleStatesFromBoardPiece(gs, movePos)
				if len(runningFromPawnStates) == 0 { // if king can't run away from dropped pawn
					return
				}
			}
		}
	}

	newGameState := tryGenerateGameStateAfterChangeAtWithNextMovesCheck(newBoard, curPlayer, nextPlayer, dropPos)
	if newGameState != nil {
		gss = append(gss, *newGameState)
	}
	return
}

func generatePossibleStatesFromBoardPiece(gs *gamestate.GameState, piecePos board.Position) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	var curBoard = gs.Board

	var movesPositions []board.Position
	if curBoard.At(piecePos).Type.ImportantPiece {
		movesPositions = curBoard.GetKingPossibleMoves(piecePos)
	} else {
		movesPositions = curBoard.GetPiecePossibleMoves(piecePos, true)
	}

	for _, movePos := range movesPositions {
		gss = append(gss, generatePossibleStatesWithMove(gs, piecePos, movePos)...)
	}
	return
}

func generatePossibleStatesFromBoardPieces(gs *gamestate.GameState) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	var curBoard = gs.Board
	curBoard.IterateBoardPieces(gs.CurMovePlayer, func(piece *board.Piece, pos board.Position) {
		gss = append(gss, generatePossibleStatesFromBoardPiece(gs, pos)...)
	})
	return
}

func generatePossibleStatesFromInventoryPieces(gs *gamestate.GameState) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer

	if curBoard.Inventories[curPlayer].IsEmpty() {
		return
	}

	var dropsPositions = curBoard.GetPossibleDropsCoords()
	for _, dropPos := range dropsPositions {
		for _, piece := range curBoard.Inventories[gs.CurMovePlayer].Pieces() {
			gss = append(gss, generatePossibleStatesWithDrop(gs, piece, dropPos)...)
		}
	}
	return
}

func generatePossibleStatesFromDefendingKing(gs *gamestate.GameState) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer

	var kingPosition = curBoard.GetKingPosition(curPlayer)
	var kingFile, kingRank = kingPosition.Get()
	var king = curBoard.Cells[kingFile][kingRank]

	var attackerPlayer = king.GetAttackerPlayer()
	var attackers = curBoard.GetPositionsOfAttackersOnCell(attackerPlayer, kingPosition)

	if len(attackers) == 1 { // can eat attacker or drop(move) piece on his path
		attackerPos := attackers[0]
		attackerPoint := graphics.NewPoint(attackerPos.File, attackerPos.Rank)
		ipPoint := graphics.NewPoint(kingFile, kingRank)
		attackPath := graphics.GetLinePoints(attackerPoint, ipPoint) // TODO: будет неправильно работать для коня, исправить

		// eating and closing from
		curBoard.IterateBoardPieces(curPlayer, func(piece *board.Piece, pos board.Position) {
			// this condition need to not creating two same game states
			// (because important piece will eat attacker when it starts to run away (see below))
			// + to not create game state where important piece closing by itself
			if !piece.Type.ImportantPiece {
				var movesPositions = curBoard.GetPiecePossibleMoves(pos, true)

				// eating
				idx := slices.Index(movesPositions, attackerPos)
				if idx != -1 {
					gss = append(gss, generatePossibleStatesWithMove(gs, pos, attackerPos)...)
				}

				// closing from
				for i := range movesPositions {
					moveFile, moveRank := movesPositions[i].Get()
					for j := 1; j < len(attackPath)-1; j++ {
						closeFile, closeRank := attackPath[j].Coordinates()
						if closeFile == moveFile && closeRank == moveRank {
							gss = append(gss, generatePossibleStatesWithMove(gs, pos, board.NewPos(closeFile, closeRank))...)
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
					gss = append(gss, generatePossibleStatesWithDrop(gs, pieceType, board.NewPos(dropFile, dropRank))...)
				}
			}
		}
	}

	// important piece is running away
	gss = append(gss, generatePossibleStatesFromBoardPiece(gs, kingPosition)...)
	return
}

func GeneratePossibleStates(gs *gamestate.GameState) (gss []gamestate.GameState) {
	gss = []gamestate.GameState{}

	if !gs.KingUnderAttack {
		gss = append(gss, generatePossibleStatesFromBoardPieces(gs)...)
		gss = append(gss, generatePossibleStatesFromInventoryPieces(gs)...)
	} else {
		gss = append(gss, generatePossibleStatesFromDefendingKing(gs)...)
	}

	return
}
