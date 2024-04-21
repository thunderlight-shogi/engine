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

type engine struct {
	gamestate *gamestate.GameState
	typemap   map[uint]*model.PieceType
}

type Engine = *engine

func (engine Engine) GetState() *gamestate.GameState {
	return engine.gamestate
}

func (engine Engine) GetReport() string {
	return evaluator.Evaluation_report(engine.gamestate)
}

func Start(id uint) (new_eng Engine, err error) {
	db := model.GetDB()
	var pos model.Preset

	err = nil

	result := db.Preload("Pieces").
		Preload("Pieces.PieceType").
		Preload("Pieces.PieceType.PromotePiece").
		Preload("Pieces.PieceType.Moves").
		Preload("Pieces.PieceType.PromotePiece.Moves").
		First(&pos, id)

	if result.Error != nil {
		err = result.Error
		return
	}

	new_eng = new(engine)

	newBoard := board.Construct()
	new_eng.typemap = make(map[uint]*model.PieceType)

	for _, piece := range pos.Pieces {
		pt := piece.PieceType

		new_eng.typemap[pt.Id] = pt

		if pt.PromotePiece != nil {
			new_eng.typemap[pt.PromotePiece.Id] = pt.PromotePiece
			pt.PromotePiece.DemotePiece = pt
		}

		newBoard.Cells[piece.File][piece.Rank] =
			&board.Piece{
				Type:   pt,
				Player: piece.Player,
			}
	}

	new_eng.gamestate = gamestate.NewGameState(newBoard, model.Sente)

	return
}

func (engine Engine) Move(move board.Move) error {
	switch move.MoveType {
	case board.Attacking, board.Moving:
		engine.gamestate.Board.MakeMove(move.OldCoords, move.NewCoords, false)

	case board.PromotionAttacking, board.PromotionMoving:
		engine.gamestate.Board.MakeMove(move.OldCoords, move.NewCoords, true)

	case board.Dropping:
		engine.gamestate.Board.MakeDrop(move.PieceType, engine.gamestate.CurMovePlayer, move.NewCoords)

	default:
		return ErrUnknownMoveType
	}

	engine.gamestate.CurMovePlayer = engine.gamestate.GetNextPlayer()
	engine.gamestate.KingUnderAttack = engine.gamestate.Board.IsKingAttacked(engine.gamestate.CurMovePlayer)

	return nil
}

func (engine Engine) FindPiece(id uint) (*model.PieceType, error) {
	t, found := engine.typemap[id]

	if found {
		return t, nil
	} else {
		return nil, ErrUnknownPieceType
	}
}

func (engine Engine) GetEngineMove() board.Move {
	return movepicker.Search(engine.gamestate)
}
