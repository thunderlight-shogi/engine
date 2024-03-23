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

func Search() PickedMove {
	return PickedMove{}
}
