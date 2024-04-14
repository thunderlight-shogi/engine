package evaluator

import (
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
	"github.com/thunderlight-shogi/engine/internal/model"
)

func iterateInventory(
	gameState *movegen.GameState,
	player model.Player,
	callback func(piece *model.PieceType),
) {
	/*
		Iterates over player's inventory in current game state
		and for each figure calls callback function
	*/
	playerInventory := gameState.Board.Inventories[player]
	for _, pieceType := range playerInventory.Pieces() {
		count := playerInventory.CountPiece(pieceType)
		for range count {
			callback(pieceType)
		}
	}
}

func iterateBoardPieces(
	gameState *movegen.GameState,
	player model.Player,
	callback func(piece *board.Piece, x int, y int),
) {
	/*
		Iterates over player's pieces on board in current game state
		and for each figure calls callback function
	*/
	for x, column := range gameState.Board.Cells {
		for y, cell := range column {
			if cell != nil && cell.Player == player { // If player's cell
				callback(cell, x, y)
			}
		}
	}
}

func iterateEmptyCells(
	gameState *movegen.GameState,
	callback func(x int, y int),
) {
	/*
		Iterates over all empty cells on board
		and for each calls callback function
	*/
	for x, column := range gameState.Board.Cells {
		for y, cell := range column {
			if cell == nil {
				callback(x, y)
			}
		}
	}
}

func createAttackMatrix(
	gameState *movegen.GameState,
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

	iterateBoardPieces(gameState, player, func(piece *board.Piece, x, y int) {
		movesCoords := gameState.Board.GetPossibleMovesCoords(x, y)
		for _, move := range movesCoords {
			moveX, moveY := move[0], move[1]
			attackCounts[moveX][moveY] += 1
		}
	})

	return attackCounts
}

func createDefendMatrix(
	gameState *movegen.GameState,
	player model.Player,
) [][]uint {
	/*
		Creates matrix same size of board
		where each cell contains number of player's figures that defend this cell
		which has player's figure in it
	*/

	// TODO: Сделать метод. Хер знает как, щас лень придумывать
	defendCounts := make([][]uint, 9)
	for i := range defendCounts {
		defendCounts[i] = make([]uint, 9)
	}
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
