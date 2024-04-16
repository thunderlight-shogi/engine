package gamestate

import (
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

type GameState struct {
	Board           board.Board
	CurMovePlayer   model.Player
	KingUnderAttack bool
}

func (gs *GameState) GetNextPlayer() model.Player {
	if gs.CurMovePlayer == model.Sente {
		return model.Gote
	} else {
		return model.Sente
	}
}
