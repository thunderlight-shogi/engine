package engine

import (
	"errors"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/evaluator"
	"github.com/thunderlight-shogi/engine/internal/engine/gamestate"
	"github.com/thunderlight-shogi/engine/internal/engine/movepicker"
	"github.com/thunderlight-shogi/engine/internal/model"
)

var ErrUnknownMoveType error = errors.New("unknown move type")
var ErrUnknownPieceType error = errors.New("unknown piece type")

var global_state gamestate.GameState

var global_type_map map[uint]*model.PieceType

func GetState() *gamestate.GameState {
	return &global_state
}

func GetReport() string {
	return evaluator.Evaluation_report(&global_state)
}

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

	newBoard := board.Construct()
	global_type_map = make(map[uint]*model.PieceType)

	for _, piece := range pos.Pieces {
		pt := piece.PieceType

		global_type_map[pt.Id] = pt

		if pt.PromotePiece != nil {
			global_type_map[pt.PromotePiece.Id] = pt.PromotePiece
			pt.PromotePiece.DemotePiece = pt
		}

		newBoard.Cells[piece.File][piece.Rank] =
			&board.Piece{
				Type:   pt,
				Player: piece.Player,
			}
	}

	global_state = *gamestate.NewGameState(newBoard, model.Sente)

	return nil
}

func Move(move board.Move) error {
	switch move.MoveType {
	case board.Attacking:
	case board.Moving:
		global_state.Board.MakeMove(move.OldCoords, move.NewCoords, false)

	case board.PromotionAttacking:
	case board.PromotionMoving:
		global_state.Board.MakeMove(move.OldCoords, move.NewCoords, true)

	case board.Dropping:
		global_state.Board.MakeDrop(move.PieceType, global_state.CurMovePlayer, move.NewCoords)

	default:
		return ErrUnknownMoveType
	}

	global_state.CurMovePlayer = global_state.GetNextPlayer()
	global_state.KingUnderAttack = global_state.Board.IsKingAttacked(global_state.CurMovePlayer)

	return nil
}

func FindPiece(id uint) (*model.PieceType, error) {
	t, found := global_type_map[id]

	if found {
		return t, nil
	} else {
		return nil, ErrUnknownPieceType
	}
}

func GetEngineMove() board.Move {
	return movepicker.Search(&global_state)
}
