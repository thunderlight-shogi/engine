package movegen

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

//TODO: Везде поменять vertical и horizontal на file и rank

type GameState struct {
	Board                     board.Board
	CurMovePlayer             model.Player
	ImportantPieceUnderAttack bool //TODO
	NumOfAttackableCells      uint //TODO: Сделать подсчет
	NumOfDropableCells        uint //TODO: Сделать подсчет
}

func (gs *GameState) getNextPlayer() model.Player {
	if gs.CurMovePlayer == model.Sente {
		return model.Gote
	} else {
		return model.Sente
	}
}

func (gs *GameState) getPossibleStatesFromBoardPiece(verticalCoord int, horizontalCoord int) (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var nextPlayer = gs.getNextPlayer()
	var movesCoords [][2]int = curBoard.GetPossibleMoves(verticalCoord, horizontalCoord)
	for _, coords := range movesCoords {
		var newBoard = curBoard.Clone()
		var cellMoveTo = newBoard.Cells[coords[0]][coords[1]]

		var emptyCell = cellMoveTo == nil
		if !emptyCell {
			newBoard.Inventories[gs.CurMovePlayer].AddPiece(cellMoveTo)
		}

		// creating alternative gamestate for promotion of piece
		var boardPiece = newBoard.Cells[verticalCoord][horizontalCoord]
		if boardPiece.IsPromotable() {
			var inPromotionZone = slices.Contains(curBoard.GetPromotionZone(boardPiece.Player), coords[1])
			var fromPromotionZone = slices.Contains(curBoard.GetPromotionZone(boardPiece.Player), horizontalCoord)
			if inPromotionZone || fromPromotionZone {
				altNewBoard := newBoard.Clone()
				altNewBoard.Cells[coords[0]][coords[1]] = &board.Piece{Type: *altNewBoard.Cells[verticalCoord][horizontalCoord].Type.PromotePiece, Player: gs.CurMovePlayer}
				altNewBoard.Cells[verticalCoord][horizontalCoord] = nil

				newGameState := GameState{Board: altNewBoard, CurMovePlayer: nextPlayer}
				gss = append(gss, newGameState)
			}
		}

		newBoard.Cells[coords[0]][coords[1]] = newBoard.Cells[verticalCoord][horizontalCoord]
		newBoard.Cells[verticalCoord][horizontalCoord] = nil

		// e. g., pawn cannot move to last row because it will have no moves to board field (so this is not possible game state)
		if len(newBoard.GetInBoardFieldMoves(coords[0], coords[1])) != 0 {
			newGameState := GameState{Board: newBoard, CurMovePlayer: nextPlayer}
			gss = append(gss, newGameState)
		}
	}
	return
}

func (gs *GameState) getPossibleStatesFromBoardPieces() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	for v, verticalCells := range curBoard.Cells {
		for h, cell := range verticalCells {
			if cell == nil { // empty cell
				continue
			}

			if cell.Player == gs.CurMovePlayer {
				gss = append(gss, gs.getPossibleStatesFromBoardPiece(v, h)...)
			}
		}
	}
	return
}

func (gs *GameState) getPossibleStatesFromInventoryPieces() (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var curPlayer = gs.CurMovePlayer
	var nextPlayer = gs.getNextPlayer()

	for _, pieceType := range curBoard.Inventories[gs.CurMovePlayer].Pieces() {
		var dropsCoords = curBoard.GetPossibleDrops()
		for _, coords := range dropsCoords {
			newBoard := curBoard.Clone()
			newBoard.Cells[coords[0]][coords[1]] = newBoard.Inventories[curPlayer].ExtractPieceToPlayer(pieceType, curPlayer)

			// e. g., pawn cannot move to last row because it will have no moves to board field (so this is not possible game state)
			if len(newBoard.GetInBoardFieldMoves(coords[0], coords[1])) != 0 {
				newGameState := GameState{Board: newBoard, CurMovePlayer: nextPlayer}
				gss = append(gss, newGameState)
			}
		}
	}
	return
}

func (gs *GameState) GetPossibleStates() (gss []GameState) {
	gss = []GameState{}

	gss = append(gss, gs.getPossibleStatesFromBoardPieces()...)
	gss = append(gss, gs.getPossibleStatesFromInventoryPieces()...)

	return
}
