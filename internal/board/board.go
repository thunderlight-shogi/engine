package board

import (
	"github.com/thunderlight-shogi/engine/internal/model"
)

type board struct {
	Cells       [][]*Piece // [Horizontal offset][Vertical offset]
	Inventories map[model.Player]Inventory
}

type Board = *board

type Piece struct {
	Type   model.PieceType
	Player model.Player
}

type inventory struct {
	pieces []*Piece
}

type Inventory = *inventory

func NewInventory() Inventory {
	return &inventory{pieces: make([]*Piece, 0)}
}

func (this_inv Inventory) Pieces() []*Piece {
	if len(this_inv.pieces) == 0 {
		return nil
	}
	return this_inv.pieces
}

func (this_inv Inventory) AddPiece(piece *Piece) {
	this_inv.pieces = append(this_inv.pieces, piece)
}

// deleting element with index i from array
func (this_inv Inventory) removePiece(i int) {
	this_inv.pieces[i] = this_inv.pieces[len(this_inv.pieces)-1]
	this_inv.pieces = this_inv.pieces[:len(this_inv.pieces)-1]
}

func (this_inv Inventory) ExtractPiece(piece *Piece) *Piece {
	for i, elem := range this_inv.pieces {
		if elem.Type.Name == piece.Type.Name {
			save := elem
			this_inv.removePiece(i)
			return save
		}
	}
	return nil
}

func (this_board Board) Clone() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Piece, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Piece, 9)
		copy(newby.Cells[i], this_board.Cells[i])
	}
	return
}

func Construct() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Piece, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Piece, 9)
	}

	newby.Inventories = make(map[model.Player]Inventory)
	newby.Inventories[model.Sente] = NewInventory()
	newby.Inventories[model.Gote] = NewInventory()

	return
}
