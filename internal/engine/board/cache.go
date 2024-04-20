package board

import "slices"

func (this_board Board) CachePossibleMoves() {
	this_board.IterateAllBoardPieces(func(piece *Piece, pos Position) {
		this_board.reachableCellsCache[pos.File][pos.Rank] = this_board.GetPieceReachableCells(pos, false)
		this_board.possibleMovesCache[pos.File][pos.Rank] = this_board.GetPiecePossibleMoves(pos, false)
	})
}

func (this_board Board) updateCacheAfterDrop(dropPosToRecache Position) {
	var recachePiece = this_board.At(dropPosToRecache)
	this_board.reachableCellsCache[dropPosToRecache.File][dropPosToRecache.Rank] = this_board.GetPieceReachableCells(dropPosToRecache, false)
	this_board.possibleMovesCache[dropPosToRecache.File][dropPosToRecache.Rank] = this_board.GetPiecePossibleMoves(dropPosToRecache, false)
	this_board.IterateAllBoardPieces(func(piece *Piece, pos Position) {
		if piece == recachePiece {
			return
		}

		var pieceReachableCellsCache = this_board.reachableCellsCache[pos.File][pos.Rank]
		idx := slices.Index(pieceReachableCellsCache, dropPosToRecache)
		if idx != -1 {
			this_board.reachableCellsCache[pos.File][pos.Rank] = this_board.GetPieceReachableCells(pos, false)
			this_board.possibleMovesCache[pos.File][pos.Rank] = this_board.GetPiecePossibleMoves(pos, false)
		}
	})
}

func (this_board Board) updateCacheAfterMove(oldPos, newPos Position) {
	var cellNew = this_board.At(newPos)
	this_board.reachableCellsCache[oldPos.File][oldPos.Rank] = nil
	this_board.possibleMovesCache[oldPos.File][oldPos.Rank] = nil

	this_board.reachableCellsCache[newPos.File][newPos.Rank] = this_board.GetPieceReachableCells(newPos, false)
	this_board.possibleMovesCache[newPos.File][newPos.Rank] = this_board.GetPiecePossibleMoves(newPos, false)

	this_board.IterateAllBoardPieces(func(piece *Piece, pos Position) {
		if piece == cellNew {
			return
		}

		var pieceReachableCellsCache = this_board.reachableCellsCache[pos.File][pos.Rank]
		idx1 := slices.Index(pieceReachableCellsCache, oldPos)
		idx2 := slices.Index(pieceReachableCellsCache, newPos)
		if idx1 != -1 || idx2 != -1 {
			this_board.reachableCellsCache[pos.File][pos.Rank] = this_board.GetPieceReachableCells(pos, false)
			this_board.possibleMovesCache[pos.File][pos.Rank] = this_board.GetPiecePossibleMoves(pos, false)
		}
	})
}
