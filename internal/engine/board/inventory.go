package board

import (
	"github.com/thunderlight-shogi/engine/internal/model"
)

type inventory struct {
	pieces map[*model.PieceType]uint
}

type Inventory = *inventory

func newInventory() Inventory {
	return &inventory{pieces: make(map[*model.PieceType]uint)}
}

func (this_inv Inventory) clone() (inv Inventory) {
	inv = newInventory()
	for k, v := range this_inv.pieces {
		inv.pieces[k] = v
	}
	return
}

func (this_inv Inventory) Pieces() []*model.PieceType {
	if len(this_inv.pieces) == 0 {
		return nil
	}

	var pieceTypes []*model.PieceType
	for key := range this_inv.pieces {
		pieceTypes = append(pieceTypes, key)
	}
	return pieceTypes
}

func (this_inv Inventory) IsEmpty() bool {
	return len(this_inv.pieces) == 0
}

func (this_inv Inventory) incrementPieceTypeNum(pt *model.PieceType) {
	var foundPieceType *model.PieceType = nil
	for k := range this_inv.pieces {
		if k.Id == pt.Id {
			foundPieceType = k
		}
	}

	if foundPieceType == nil {
		this_inv.pieces[pt] = 1
	} else {
		this_inv.pieces[foundPieceType]++
	}
}

// returns true if piece type is found, and false otherwise
func (this_inv Inventory) decrementPieceTypeNum(pt *model.PieceType) bool {
	var foundPieceType *model.PieceType = nil
	for k := range this_inv.pieces {
		if k.Id == pt.Id {
			foundPieceType = k
		}
	}

	if foundPieceType != nil {
		var num = this_inv.pieces[foundPieceType]
		if num == 1 {
			delete(this_inv.pieces, foundPieceType)
		} else {
			this_inv.pieces[foundPieceType]--
		}
		return true
	}
	return false
}

func (this_inv Inventory) AddPiece(piece *Piece) {
	var addedPiece *model.PieceType
	if piece.IsPromoted() {
		addedPiece = piece.Type.DemotePiece
	} else {
		addedPiece = &piece.Type
	}
	this_inv.incrementPieceTypeNum(addedPiece)
}

// maybe replace *pieceType with pieceType name
func (this_inv Inventory) ExtractPieceToPlayer(pieceType *model.PieceType, player model.Player) *Piece {
	var found = this_inv.decrementPieceTypeNum(pieceType)
	if found {
		return &Piece{Type: *pieceType, Player: player}
	} else {
		return nil
	}
}

func (this_inv Inventory) CountPiece(pieceType *model.PieceType) uint {
	count, found := this_inv.pieces[pieceType]
	if found {
		return count
	} else {
		return 0
	}
}
