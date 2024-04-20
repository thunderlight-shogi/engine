package board

import "github.com/thunderlight-shogi/engine/internal/model"

func (this_board Board) IterateInventory(
	player model.Player,
	callback func(piece *model.PieceType),
) {
	/*
		Iterates over player's inventory
		and for each figure calls callback function
	*/
	playerInventory := this_board.Inventories[player]
	for _, pieceType := range playerInventory.Pieces() {
		count := playerInventory.CountPiece(pieceType)
		for range count {
			callback(pieceType)
		}
	}
}

func (this_board Board) IterateAllBoardPieces(
	callback func(piece *Piece, pos Position),
) {
	/*
		Iterates over all pieces on board
		and for each figure calls callback function
	*/
	for x := range this_board.Cells {
		for y, cell := range this_board.Cells[x] {
			if cell != nil {
				callback(cell, NewPos(x, y))
			}
		}
	}
}

func (this_board Board) IterateBoardPieces(
	player model.Player,
	callback func(piece *Piece, pos Position),
) {
	/*
		Iterates over player's pieces on board
		and for each figure calls callback function
	*/
	this_board.IterateAllBoardPieces(func(piece *Piece, pos Position) {
		if piece.Player == player { // If player's piece
			callback(piece, pos)
		}
	})
}

func (this_board Board) IterateBoardPiecesWithEarlyExit(
	player model.Player,
	callback func(piece *Piece, pos Position) bool,
) {
	/*
		Iterates over player's pieces on board
		and for each figure calls callback function.
		If callback return true then iterating finishes
	*/
	for x := range this_board.Cells {
		for y, cell := range this_board.Cells[x] {
			if cell != nil && cell.Player == player {
				if callback(cell, NewPos(x, y)) {
					return
				}
			}
		}
	}
}

func (this_board Board) IterateEmptyCells(
	callback func(pos Position),
) {
	/*
		Iterates over all empty cells on board
		and for each calls callback function
	*/
	for x := range this_board.Cells {
		for y, cell := range this_board.Cells[x] {
			if cell == nil {
				callback(NewPos(x, y))
			}
		}
	}
}
