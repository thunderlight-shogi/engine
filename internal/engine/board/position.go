package board

import "github.com/thunderlight-shogi/engine/internal/model"

type Position struct {
	File int `json:"file"`
	Rank int `json:"rank"`
}

func NewPos(file, rank int) Position {
	return Position{File: file, Rank: rank}
}

func (pos Position) Get() (int, int) {
	return pos.File, pos.Rank
}

type MoveType uint

const (
	PromotionAttacking MoveType = iota
	PromotionMoving
	Attacking
	Moving
	Dropping
	Surrender
)

type Move struct {
	OldCoords Position         `json:"old_pos"`
	NewCoords Position         `json:"new_pos"`
	PieceType *model.PieceType `json:"piece_type"`
	MoveType  MoveType         `json:"move_type"`
}
