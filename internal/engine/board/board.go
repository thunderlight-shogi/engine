package board

import (
	"fmt"

	"github.com/thunderlight-shogi/engine/internal/model"
)

type Piece struct {
	Type   model.PieceType
	Player model.Player
}

type inventory struct {
	pieces []*Piece
}

type board struct {
	Cells       [][]*Piece // [Horizontal offset][Vertical offset]
	Inventories map[model.Player]Inventory
}

type Inventory = *inventory
type Board = *board

func newInventory() Inventory {
	return &inventory{pieces: make([]*Piece, 0)}
}

func (this_inv Inventory) clone() (inv Inventory) {
	inv = new(inventory)
	inv.pieces = make([]*Piece, len(this_inv.pieces))
	copy(inv.pieces, this_inv.pieces)
	return
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

// maybe replace *piece with pieceType name
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

func Construct() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Piece, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Piece, 9)
	}

	newby.Inventories = make(map[model.Player]Inventory)
	newby.Inventories[model.Sente] = newInventory()
	newby.Inventories[model.Gote] = newInventory()

	return
}

func (this_board Board) Clone() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Piece, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Piece, 9)
		copy(newby.Cells[i], this_board.Cells[i])
	}

	newby.Inventories = make(map[model.Player]Inventory)
	//fmt.Printf("this_board.Inventories[model.Sente].pieces: %v\n", this_board.Inventories[model.Sente].pieces)
	newby.Inventories[model.Sente] = this_board.Inventories[model.Sente].clone()
	//fmt.Printf("newby.Inventories[model.Sente].pieces: %v\n", newby.Inventories[model.Sente].pieces)
	newby.Inventories[model.Gote] = this_board.Inventories[model.Gote].clone()

	return
}

// For test purposes
func (this_board Board) Print() {
	fmt.Println("-------------")
	for _, piece := range this_board.Inventories[model.Sente].pieces {
		fmt.Print(string(piece.Type.Name[0]) + " ")
	}
	fmt.Println()
	for _, verticalPieces := range this_board.Cells {
		for _, piece := range verticalPieces {
			if piece == nil {
				fmt.Print("-")
			} else {
				fmt.Print(string(piece.Type.Name[0]))
			}
		}
		fmt.Println()
	}
	for _, piece := range this_board.Inventories[model.Gote].pieces {
		fmt.Print(string(piece.Type.Name[0]))
	}
	fmt.Println("\n-------------")
}
