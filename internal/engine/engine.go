package engine

import (
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
	"github.com/thunderlight-shogi/engine/internal/model"
)

var global_state movegen.GameState

func Start(id uint) error {
	db := model.GetDB()
	var pos model.Preset

	result := db.Preload("Pieces").
		Preload("Pieces.PieceType").
		Preload("Pieces.PieceType.PromotePiece").
		First(&pos, id)
	if result.Error != nil {
		return result.Error
	}

	global_state.Board = board.Construct()

	for _, piece := range pos.Pieces {
		pt := piece.PieceType

		if pt.PromotePiece != nil {
			pt.PromotePiece.DemotePiece = pt
		}

		global_state.Board.Cells[piece.HorizontalOffset][piece.VerticalOffset] =
			&board.Piece{
				Type:   *pt,
				Player: piece.Player,
			}
	}

	global_state.CurMovePlayer = model.Sente

	return nil
}
