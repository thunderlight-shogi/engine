package movepicker

import (
	"github.com/thunderlight-shogi/engine/internal/board"
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

func Search(currentBoard board.Board) (pickedMove PickedMove) {
	var found bool = false
	var curPiece *board.Piece
	for horiz := 0; horiz < len(currentBoard.Cells); horiz++ {
		for vert := 0; vert < len(currentBoard.Cells[horiz]); vert++ {
			if currentBoard.Cells[horiz][vert] != nil {
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
