package engine

import (
	"errors"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
	"github.com/thunderlight-shogi/engine/internal/engine/movepicker"
	"github.com/thunderlight-shogi/engine/internal/model"
)

var ErrUnknownMoveType error = errors.New("unknown move type")

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

		global_state.Board.Cells[piece.File-1][piece.Rank-1] =
			&board.Piece{
				Type:   pt,
				Player: piece.Player,
			}
	}

	global_state.CurMovePlayer = model.Sente

	return nil
}

func Move(move board.Move) error {
	piece_from := global_state.Board.At(move.OldCoords)
	piece_to := global_state.Board.At(move.NewCoords)

	switch move.MoveType {
	case board.Attacking:
		global_state.Board.Inventories[global_state.CurMovePlayer].AddPiece(piece_to)
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from)

	case board.PromotionAttacking:
		global_state.Board.Inventories[global_state.CurMovePlayer].AddPiece(piece_to)
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from.GetPromotedPiece())

	case board.Dropping:
		global_state.Board.Inventories[global_state.CurMovePlayer].AddPiece(piece_to)
		new_piece := global_state.Board.Inventories[global_state.CurMovePlayer].
			ExtractPieceToPlayer(move.PieceType, global_state.CurMovePlayer)

		global_state.Board.Set(move.NewCoords, new_piece)

	case board.Moving:
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from)

	case board.PromotionMoving:
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from.GetPromotedPiece())
	default:
		return ErrUnknownMoveType
	}

	switch global_state.CurMovePlayer {
	case model.Sente:
		global_state.CurMovePlayer = model.Gote
	case model.Gote:
		global_state.CurMovePlayer = model.Sente
	}

	return nil
}

func GetMove() (move board.Move) {
	Move(move)
	move = movepicker.Search(&global_state)
	return
}
