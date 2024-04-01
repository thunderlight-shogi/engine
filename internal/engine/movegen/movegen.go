package movegen

import (
	"github.com/thunderlight-shogi/engine/internal/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

type GameState struct {
	Board                     board.Board
	CurMovePlayer             model.Player
	ImportantPieceUnderAttack bool //TODO
	NumOfAttackableCells      uint //TODO: Сделать подсчет
	NumOfDropableCells        uint //TODO: Сделать подсчет
}

func (gs *GameState) getShiftSign() int {
	if gs.CurMovePlayer == model.Sente {
		return 1
	} else {
		return -1
	}
}

func (gs *GameState) getNextPlayer() model.Player {
	if gs.CurMovePlayer == model.Sente {
		return model.Gote
	} else {
		return model.Sente
	}
}

func (gs *GameState) getPossibleMoves(verticalCoord int, horizontalCoord int) (movesCoords [][]int) {
	movesCoords = make([][]int, 0)
	shiftSign := gs.getShiftSign()

	curBoard := gs.Board
	piece := curBoard.Cells[verticalCoord][horizontalCoord]
	for _, move := range piece.Type.Moves {
		coord := []int{
			verticalCoord + move.HorizontalShift*shiftSign,
			horizontalCoord + move.VerticalShift*shiftSign,
		}

		var inBoardField bool = coord[0] >= 0 && coord[0] < len(curBoard.Cells) && coord[1] >= 0 && coord[1] < len(curBoard.Cells[verticalCoord])
		if inBoardField {
			var freeOrEnemyCell bool = curBoard.Cells[coord[0]][coord[1]] == nil || curBoard.Cells[coord[0]][coord[1]].Player != gs.CurMovePlayer
			if freeOrEnemyCell {
				movesCoords = append(movesCoords, coord)
			}
		}
	}
	return
}

// TODO: Добавить PieceType в параметры (У разных фигур могут быть разные клетки сброса)
func (gs *GameState) getPossibleDrops() (dropsCoords [][]int) {
	dropsCoords = make([][]int, 0)

	curBoard := gs.Board
	for v, verticalCells := range curBoard.Cells {
		for h, cell := range verticalCells {
			if cell == nil {
				coord := []int{v, h}
				dropsCoords = append(dropsCoords, coord)
			}
		}
	}
	return
}

func (gs *GameState) getPossibleStatesFromBoardPiece(verticalCoord int, horizontalCoord int) (gss []GameState) {
	gss = make([]GameState, 0)

	var curBoard = gs.Board
	var nextPlayer = gs.getNextPlayer()
	var movesCoords [][]int = gs.getPossibleMoves(verticalCoord, horizontalCoord)
	for _, coords := range movesCoords {
		newBoard := curBoard.Clone()
		if newBoard.Cells[coords[0]][coords[1]] != nil {
			//TODO: Если была срублена перевернутая фигура, добавлять неперевернутую
			newBoard.Inventories[gs.CurMovePlayer].AddPiece(newBoard.Cells[coords[0]][coords[1]])
		}
		newBoard.Cells[coords[0]][coords[1]] = newBoard.Cells[verticalCoord][horizontalCoord]
		newBoard.Cells[verticalCoord][horizontalCoord] = nil

		newGameState := GameState{Board: newBoard, CurMovePlayer: nextPlayer}
		gss = append(gss, newGameState)
	}
	return
}

func (gs *GameState) getPossibleStatesFromBoardPieces() (gss []GameState) {
	gss = make([]GameState, 0)

	var curBoard = gs.Board
	for v, verticalPieces := range curBoard.Cells {
		for h, piece := range verticalPieces {
			if piece == nil {
				continue
			}

			if piece.Player == gs.CurMovePlayer {
				gss = append(gss, gs.getPossibleStatesFromBoardPiece(v, h)...)
			}
		}
	}
	return
}

func (gs *GameState) getPossibleStatesFromInventoryPieces() (gss []GameState) {
	gss = make([]GameState, 0)

	var curBoard = gs.Board
	var nextPlayer = gs.getNextPlayer()
	for _, piece := range curBoard.Inventories[gs.CurMovePlayer].Pieces() {
		dropsCoords := gs.getPossibleDrops()
		for _, coords := range dropsCoords {
			newBoard := curBoard.Clone()
			newBoard.Cells[coords[0]][coords[1]] = newBoard.Inventories[gs.CurMovePlayer].ExtractPiece(piece)

			newGameState := GameState{Board: newBoard, CurMovePlayer: nextPlayer}
			gss = append(gss, newGameState)
		}
	}
	return
}

func (gs *GameState) GetPossibleStates() (gss []GameState) {
	gss = make([]GameState, 0)

	gss = append(gss, gs.getPossibleStatesFromBoardPieces()...)
	gss = append(gss, gs.getPossibleStatesFromInventoryPieces()...)

	return
}
