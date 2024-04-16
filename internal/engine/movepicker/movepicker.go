package movepicker

import (
	"math"
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/evaluator"
	"github.com/thunderlight-shogi/engine/internal/engine/gamestate"
	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
	"github.com/thunderlight-shogi/engine/internal/model"
)

const DEPTH = 3

func getMoveFromBoardDifference(baseGs *gamestate.GameState, newGs *gamestate.GameState) (pickedMove board.Move) {
	var baseBoard = baseGs.Board
	var newBoard = newGs.Board
	pickedMove.OldCoords = board.NewPos(-1, -1)

	for horizontal := 0; horizontal < 9; horizontal++ {
		for vertical := 0; vertical < 9; vertical++ {
			if newBoard.Cells[horizontal][vertical] != baseBoard.Cells[horizontal][vertical] {
				if newBoard.Cells[horizontal][vertical] != nil {
					pickedMove.NewCoords = board.NewPos(horizontal, vertical)
				} else {
					pickedMove.OldCoords = board.NewPos(horizontal, vertical)
				}

			}
		}
	}
	//if there was NO cell that WAS EMPTY before move we
	//check what piece player have dropped

	if pickedMove.OldCoords.GetFile() == -1 {
		pickedMove.MoveType = board.Dropping
		var baseInventory = baseBoard.Inventories[baseGs.CurMovePlayer]
		var newInventory = newBoard.Inventories[baseGs.CurMovePlayer]
		for _, pieceType := range baseInventory.Pieces() {
			if newInventory.CountPiece(pieceType) != baseInventory.CountPiece(pieceType) {
				pickedMove.PieceType = pieceType
				break
			}
		}
	} else {
		//check for move type
		var file = pickedMove.NewCoords.GetFile()
		var rank = pickedMove.NewCoords.GetRank()
		var oldType = baseBoard.Cells[pickedMove.OldCoords.GetFile()][pickedMove.OldCoords.GetRank()].Type
		var newType = newBoard.Cells[file][rank].Type
		pickedMove.PieceType = oldType
		if baseBoard.Cells[file][rank] != nil {
			if oldType.Name != newType.Name {
				pickedMove.MoveType = board.PromotionAttacking
			} else {
				pickedMove.MoveType = board.Attacking
			}
		} else {
			if oldType.Name != newType.Name {
				pickedMove.MoveType = board.PromotionMoving
			} else {
				pickedMove.MoveType = board.Moving
			}
		}
	}
	return pickedMove
}

func SortGameStates(gss []gamestate.GameState, baseBoard *gamestate.GameState) {
	slices.SortFunc(gss, func(i, j gamestate.GameState) int {
		var firstMoveType = int(getMoveFromBoardDifference(baseBoard, &i).MoveType)
		var secondMoveType = int(getMoveFromBoardDifference(baseBoard, &j).MoveType)
		if firstMoveType > secondMoveType {
			return 1
		} else if firstMoveType < secondMoveType {
			return -1
		}
		return 0
	})
}

func minimax(gs *gamestate.GameState, depth int, maximizingPlayer bool) float32 {
	if depth == 0 {
		return evaluator.Evaluate(gs)
	}
	if maximizingPlayer {
		var value float32 = -math.MaxFloat32
		var allBoards = movegen.GeneratePossibleStates(gs)
		for idx := range allBoards {
			value = max(value, minimax(&allBoards[idx], depth-1, false))
		}
		return value
	} else {
		var value float32 = math.MaxFloat32
		var allBoards = movegen.GeneratePossibleStates(gs)
		for idx := range allBoards {
			value = min(value, minimax(&allBoards[idx], depth-1, true))
		}
		return value
	}
}

func alphabeta(gs *gamestate.GameState, depth int, a *float32, b *float32, maximizingPlayer bool) float32 {
	if depth == 0 {
		return evaluator.Evaluate(gs)
	}
	var allBoards = movegen.GeneratePossibleStates(gs)
	SortGameStates(allBoards, gs)
	if maximizingPlayer {
		var value float32 = -math.MaxFloat32
		for idx := range allBoards {
			value = max(value, alphabeta(&allBoards[idx], depth-1, a, b, false))
			if value > *b {
				break
			}
			*a = max(*a, value)
		}
		return value
	} else {
		var value float32 = math.MaxFloat32
		for idx := range allBoards {
			value = min(value, alphabeta(&allBoards[idx], depth-1, a, b, true))
			if value < *a {
				break
			}
			*b = min(*b, value)
		}

		return value
	}
}

func Search(currentGameState *gamestate.GameState) board.Move {
	var allGs = movegen.GeneratePossibleStates(currentGameState)

	var bestValue float32 = -math.MaxFloat32
	var maximizingPlayer = true
	var a float32 = -math.MaxFloat32
	var b float32 = math.MaxFloat32
	if currentGameState.CurMovePlayer == model.Gote {
		bestValue = math.MaxFloat32
		maximizingPlayer = false
	}
	var bestIndex = 0
	SortGameStates(allGs, currentGameState)
	for index := range allGs {
		var tempValue = alphabeta(&allGs[index], DEPTH, &a, &b, !maximizingPlayer)
		if maximizingPlayer {
			if tempValue > bestValue {
				bestValue = tempValue
				bestIndex = index
			}

		} else {
			if tempValue < bestValue {
				bestValue = tempValue
				bestIndex = index
			}
		}
	}
	var bestGs = allGs[bestIndex]
	return getMoveFromBoardDifference(currentGameState, &bestGs)
}
