package gamestate

import (
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

// TODO: GameState сделать указателем

type GameState struct {
	Board           board.Board
	CurMovePlayer   model.Player
	KingUnderAttack bool
}

func NewGameState(someBoard board.Board, player model.Player) (gs *GameState) {
	gs = &GameState{}
	gs.Board = someBoard
	gs.Board.CachePossibleMoves()
	gs.CurMovePlayer = player
	gs.KingUnderAttack = gs.Board.IsKingAttacked(gs.CurMovePlayer)
	return
}

func (gs *GameState) GetNextPlayer() model.Player {
	if gs.CurMovePlayer == model.Sente {
		return model.Gote
	} else {
		return model.Sente
	}
}
