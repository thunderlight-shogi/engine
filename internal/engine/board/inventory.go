package board

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/model"
)

// TODO: Слайс заменить на мапу
type inventory struct {
	pieces []*model.PieceType
}

type Inventory = *inventory

func newInventory() Inventory {
	return &inventory{pieces: make([]*model.PieceType, 0)}
}

func (this_inv Inventory) clone() (inv Inventory) {
	inv = new(inventory)
	inv.pieces = make([]*model.PieceType, len(this_inv.pieces))
	copy(inv.pieces, this_inv.pieces)
	return
}

func (this_inv Inventory) Pieces() []*model.PieceType {
	if len(this_inv.pieces) == 0 {
		return nil
	}
	return this_inv.pieces
}

func (this_inv Inventory) IsEmpty() bool {
	return len(this_inv.pieces) == 0
}

func (this_inv Inventory) AddPiece(piece *Piece) {
	var addedPiece *model.PieceType
	if piece.IsPromoted() {
		addedPiece = piece.Type.DemotePiece
	} else {
		addedPiece = &piece.Type
	}
	this_inv.pieces = append(this_inv.pieces, addedPiece)
}

// deleting element with index i from array
func (this_inv Inventory) removePiece(i int) {
	this_inv.pieces[i] = this_inv.pieces[len(this_inv.pieces)-1]
	this_inv.pieces = this_inv.pieces[:len(this_inv.pieces)-1]
}

// maybe replace *pieceType with pieceType name
func (this_inv Inventory) ExtractPieceToPlayer(pieceType *model.PieceType, player model.Player) *Piece {
	idx := slices.Index(this_inv.pieces, pieceType)
	if idx != -1 {
		this_inv.removePiece(idx)
		return &Piece{Type: *pieceType, Player: player}
	}
	return nil
}
