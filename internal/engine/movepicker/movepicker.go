package movepicker

import (
	"github.com/thunderlight-shogi/engine/internal/board"
	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
)

type MoveType uint

const (
	Moving MoveType = iota
	Attacking
	Dropping
	PromotionMoving
	PromotionAttacking
)

type Coordinates struct {
	horizontal int
	vertical   int
}

type PickedMove struct {
	piece     *board.Piece
	newCoords Coordinates
	moveType  MoveType
}

func Search(currentGameState *movegen.GameState) (pickedMove PickedMove) {
	var found bool = false
	var currentBoard = currentGameState.Board
	var curPiece *board.Piece
	for horiz := 0; horiz < len(currentBoard.Cells); horiz++ {
		for vert := 0; vert < len(currentBoard.Cells[horiz]); vert++ {
			if currentBoard.Cells[horiz][vert] != nil && currentBoard.Cells[horiz][vert].Player == currentGameState.CurMovePlayer {
				curPiece = currentBoard.Cells[horiz][vert]
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if curPiece == nil {
		panic("Board has no pieces!")
	}
	pickedMove.piece = curPiece
	pickedMove.newCoords = Coordinates{horizontal: 4, vertical: 4}
	pickedMove.moveType = Moving
	return pickedMove
}
