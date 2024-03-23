package board

import "github.com/thunderlight-shogi/engine/internal/model"

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
	pieces map[string]int
}

type Inventory = *inventory

func NewInventory() Inventory {
	return &inventory{pieces: make(map[string]int)}
}

func (this_inv Inventory) Pieces() map[string]int {
	if len(this_inv.pieces) == 0 {
		return nil
	}
	return this_inv.pieces
}

func (this_inv Inventory) PieceNum(name string) int {
	return this_inv.pieces[name]
}

func (this_inv Inventory) AddPiece(name string) {
	_, exists := this_inv.pieces[name]
	if exists {
		this_inv.pieces[name]++
	} else {
		this_inv.pieces[name] = 1
	}
}

func (this_inv Inventory) DelPiece(name string) {
	_, exists := this_inv.pieces[name]
	if !exists {
		return
	}

	if this_inv.pieces[name] == 1 {
		delete(this_inv.pieces, name)
	} else {
		this_inv.pieces[name]--
	}
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
