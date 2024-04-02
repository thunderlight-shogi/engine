package movegen

import (
	"slices"

	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

//TODO: Везде поменять vertical и horizontal на file и rank

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

func (gs *GameState) getPromotionZone() []int {
	if gs.CurMovePlayer == model.Sente {
		return board.PromotionZoneForSente
	} else {
		return board.PromotionZoneForGote
	}
}

func (gs *GameState) canPieceReachCell(vPieceCoord int, hPieceCoord int, vCellCoord int, hCellCoord int) bool {
	var isReached bool = true

	var curBoard = gs.Board
	var piece = curBoard.Cells[vPieceCoord][hPieceCoord]
	var moves = piece.Type.Moves
	var shiftSign = gs.getShiftSign()

	origin := graphics.NewPoint(vPieceCoord, hPieceCoord)
	end := graphics.NewPoint(vCellCoord, hCellCoord)
	coordsBetween := graphics.GetLinePoints(origin, end)
	for i := 1; i < len(coordsBetween)-1; i++ { // skipping origin and end
		vMiddleCoord, hMiddleCoord := coordsBetween[i].Coordinates()
		hShift, vShift := vMiddleCoord-vPieceCoord, hMiddleCoord-hPieceCoord
		// maybe better to replace with simple loop
		idx := slices.IndexFunc(moves, func(move model.Move) bool {
			return move.HorizontalShift*shiftSign == hShift && move.VerticalShift*shiftSign == vShift
		})
		if idx != -1 { // if move between origin and end was found in moves of piece
			emptyCell := curBoard.Cells[vMiddleCoord][hMiddleCoord] == nil
			if !emptyCell {
				isReached = false
				break
			}
		}
	}
	return isReached
}

func (gs *GameState) getPossibleMoves(verticalCoord int, horizontalCoord int) (movesCoords [][2]int) {
	movesCoords = [][2]int{}

	var curBoard = gs.Board
	var shiftSign = gs.getShiftSign()
	var piece = curBoard.Cells[verticalCoord][horizontalCoord]
	var moves = piece.Type.Moves
	for _, move := range moves {
		var vMoveCoord = verticalCoord + move.HorizontalShift*shiftSign
		var hMoveCoord = horizontalCoord + move.VerticalShift*shiftSign

		var inBoardField bool = vMoveCoord >= 0 && vMoveCoord < len(curBoard.Cells) && hMoveCoord >= 0 && hMoveCoord < len(curBoard.Cells[vMoveCoord])
		if inBoardField {
			var emptyOrEnemyCell bool = curBoard.Cells[vMoveCoord][hMoveCoord] == nil || curBoard.Cells[vMoveCoord][hMoveCoord].Player != gs.CurMovePlayer
			if emptyOrEnemyCell {
				if gs.canPieceReachCell(verticalCoord, horizontalCoord, vMoveCoord, hMoveCoord) {
					movesCoords = append(movesCoords, [2]int{vMoveCoord, hMoveCoord})
				}
			}
		}
	}
	return
}

// TODO: Добавить PieceType в параметры (У разных фигур могут быть разные клетки сброса)
// TODO: Учитывать, сможет ли фигура пойти дальше (при ходе тоже)
func (gs *GameState) getPossibleDrops() (dropsCoords [][2]int) {
	dropsCoords = [][2]int{}

	var curBoard = gs.Board
	for vDropCoord, verticalCells := range curBoard.Cells {
		for hDropCoord, cell := range verticalCells {
			if cell == nil { // empty cell
				dropsCoords = append(dropsCoords, [2]int{vDropCoord, hDropCoord})
			}
		}
	}
	return
}

func (gs *GameState) getPossibleStatesFromBoardPiece(verticalCoord int, horizontalCoord int) (gss []GameState) {
	gss = []GameState{}

	var curBoard = gs.Board
	var nextPlayer = gs.getNextPlayer()
	var movesCoords [][2]int = gs.getPossibleMoves(verticalCoord, horizontalCoord)
	for _, coords := range movesCoords {
		var newBoard = curBoard.Clone()
		//создавать gamestate прям тут и позже из него вызывать какой-нибудь метод?
		var cellMoveTo = newBoard.Cells[coords[0]][coords[1]]

		var emptyCell = cellMoveTo == nil
		if !emptyCell {
			newBoard.Inventories[gs.CurMovePlayer].AddPiece(cellMoveTo)
		}

		// creating alternative gamestate for promotion of piece
		if newBoard.Cells[verticalCoord][horizontalCoord].IsPromotable() {
			var inPromotionZone = slices.Contains(gs.getPromotionZone(), coords[1])
			var fromPromotionZone = slices.Contains(gs.getPromotionZone(), horizontalCoord)
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

		newGameState := GameState{Board: newBoard, CurMovePlayer: nextPlayer}
		gss = append(gss, newGameState)
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
