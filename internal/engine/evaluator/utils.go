package evaluator

import (
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

func createAttackMatrix(
	boardVar board.Board,
	player model.Player,
) [][]uint {
	/*
		Creates matrix same size of board
		where each cell contains number of player's figures that attack this cell
	*/
	attackCounts := make([][]uint, 9)
	for i := range attackCounts {
		attackCounts[i] = make([]uint, 9)
	}

	boardVar.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		movesCoords := boardVar.GetPiecePossibleMoves(pos)
		for _, move := range movesCoords {
			moveX, moveY := move.Get()
			attackCounts[moveX][moveY] += 1
		}
	})

	return attackCounts
}

func createDefendMatrix(
	boardVar board.Board,
	player model.Player,
) [][]uint {
	/*
		Creates matrix same size of board
		where each cell contains number of player's figures that defend this cell
		which has player's figure in it
	*/

	defendCounts := make([][]uint, 9)
	for i := range defendCounts {
		defendCounts[i] = make([]uint, 9)
	}

	boardVar.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		movesCoords := boardVar.GetPieceReachableMoves(pos)
		for _, move := range movesCoords {
			moveX, moveY := move.Get()
			var moveCell = boardVar.Cells[moveX][moveY]
			if moveCell == nil {
				continue
			}
			var friendlyPiece bool = moveCell.Player == player
			if friendlyPiece {
				defendCounts[moveX][moveY] += 1
			}
		}
	})
	return defendCounts
}

func sumOf2DMatrix(mat [][]uint) uint {
	var sum uint = 0
	for _, row := range mat {
		for _, value := range row {
			sum += value
		}
	}
	return sum
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func chebyshevDistance(x1, y1, x2, y2 int) int {
	return max(abs(x1-x2), abs(y1-y2))
}
