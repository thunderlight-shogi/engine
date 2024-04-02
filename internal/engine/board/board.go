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

var PromotionZoneForSente = []int{0, 1, 2}
var PromotionZoneForGote = []int{6, 7, 8}

func (piece *Piece) IsPromoted() bool {
	return piece.Type.DemotePiece != nil
}

func (piece *Piece) IsPromotable() bool {
	return piece.Type.PromotePiece != nil
}

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
	var addedPiece *Piece
	if piece.IsPromoted() {
		addedPiece = &Piece{Type: *piece.Type.DemotePiece, Player: piece.Player}
	} else {
		addedPiece = piece
	}
	this_inv.pieces = append(this_inv.pieces, addedPiece)
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
	newby.Inventories[model.Sente] = this_board.Inventories[model.Sente].clone()
	newby.Inventories[model.Gote] = this_board.Inventories[model.Gote].clone()

	return
}

// For test purposes
func (this_board Board) Print() {
	fmt.Println("-------------")
	for _, piece := range this_board.Inventories[model.Sente].pieces {
		if piece.IsPromoted() {
			fmt.Print(string(piece.Type.Name[0]) + "+ ")
		} else {
			fmt.Print(string(piece.Type.Name[0]) + " ")
		}
	}
	fmt.Println()
	for _, verticalPieces := range this_board.Cells {
		for _, piece := range verticalPieces {
			if piece == nil {
				fmt.Print("-")
			} else if piece.IsPromoted() {
				fmt.Print(string(piece.Type.Name[0]) + "+")
			} else {
				fmt.Print(string(piece.Type.Name[0]))
			}
		}
		fmt.Println()
	}
	for _, piece := range this_board.Inventories[model.Gote].pieces {
		if piece.IsPromoted() {
			fmt.Print(string(piece.Type.Name[0]) + "+ ")
		} else {
			fmt.Print(string(piece.Type.Name[0]) + " ")
		}
	}
	fmt.Println("\n-------------")
}
