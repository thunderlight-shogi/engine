package engine

import (
	"errors"
	"slices"

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
		Preload("Pieces.PieceType.Moves").
		Preload("Pieces.PieceType.PromotePiece.Moves").
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

func isCellAttacked(this_board board.Board, attackerPlayer model.Player, cellPos board.Position) bool {
	for x, file := range this_board.Cells {
		for y, piece := range file {
			if piece != nil && piece.Player == attackerPlayer { // If player's cell
				pos := board.NewPos(x, y)

				var movesPositions = this_board.GetPieceReachableMoves(pos)
				idx := slices.Index(movesPositions, cellPos)
				if idx != -1 {
					return true
				}
			}
		}
	}
	return false
}

func Move(move board.Move) error {
	switch move.MoveType {
	case board.Attacking:
		piece_from := global_state.Board.At(move.OldCoords)
		piece_to := global_state.Board.At(move.NewCoords)
		global_state.Board.Inventories[global_state.CurMovePlayer].AddPiece(piece_to)
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from)

	case board.PromotionAttacking:
		piece_from := global_state.Board.At(move.OldCoords)
		piece_to := global_state.Board.At(move.NewCoords)
		global_state.Board.Inventories[global_state.CurMovePlayer].AddPiece(piece_to)
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from.GetPromotedPiece())

	case board.Dropping:
		new_piece := global_state.Board.Inventories[global_state.CurMovePlayer].
			ExtractPieceToPlayer(move.PieceType, global_state.CurMovePlayer)
		global_state.Board.Set(move.NewCoords, new_piece)

	case board.Moving:
		piece_from := global_state.Board.At(move.OldCoords)
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from)

	case board.PromotionMoving:
		piece_from := global_state.Board.At(move.OldCoords)
		global_state.Board.Set(move.OldCoords, nil)
		global_state.Board.Set(move.NewCoords, piece_from.GetPromotedPiece())
	default:
		return ErrUnknownMoveType
	}

	enemyKingPos := global_state.Board.GetKingPositionForPlayer(global_state.GetNextPlayer())
	global_state.KingUnderAttack = isCellAttacked(global_state.Board, global_state.CurMovePlayer, enemyKingPos)

	global_state.CurMovePlayer = global_state.GetNextPlayer()

	return nil
}

func EngineMove() (move board.Move, err error) {
	move = movepicker.Search(&global_state)
	err = Move(move)
	return
}

func GetHelp() board.Move {
	return movepicker.Search(&global_state)
}
