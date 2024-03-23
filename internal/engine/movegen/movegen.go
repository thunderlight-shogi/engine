package movegen

import (
	"github.com/thunderlight-shogi/engine/internal/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

type GameState struct {
	Board         board.Board
	CurMovePlayer model.Player
}

func (gs *GameState) GetPossibleStates() []GameState {
	return []GameState{}
}
