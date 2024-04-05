package board

import (
	"fmt"
	"slices"

	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

type Piece struct {
	Type   model.PieceType
	Player model.Player
}

type board struct {
	Cells       [][]*Piece // [Horizontal offset][Vertical offset]
	Inventories map[model.Player]Inventory
}

type Board = *board

var PromotionZoneForSente = []int{0, 1, 2}
var PromotionZoneForGote = []int{6, 7, 8}

func (this_piece *Piece) IsPromoted() bool {
	return this_piece.Type.DemotePiece != nil
}

func (this_piece *Piece) IsPromotable() bool {
	return this_piece.Type.PromotePiece != nil
}

func (this_piece *Piece) getShiftSign() int {
	if this_piece.Player == model.Sente {
		return 1
	} else {
		return -1
	}
}

func Construct() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Piece, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Piece, 9)
	}

	newby.Inventories = make(map[model.Player]Inventory)
	newby.Inventories[model.Sente] = newInventory()
	newby.Inventories[model.Gote] = newInventory()

	return
}

func (this_board Board) Clone() (newby Board) {
	newby = new(board)

	newby.Cells = make([][]*Piece, 9)
	for i := range newby.Cells {
		newby.Cells[i] = make([]*Piece, 9)
		copy(newby.Cells[i], this_board.Cells[i])
	}

	newby.Inventories = make(map[model.Player]Inventory)
	newby.Inventories[model.Sente] = this_board.Inventories[model.Sente].clone()
	newby.Inventories[model.Gote] = this_board.Inventories[model.Gote].clone()

	return
}

func (this_board Board) GetPromotionZone(player model.Player) []int {
	if player == model.Sente {
		return PromotionZoneForSente
	} else {
		return PromotionZoneForGote
	}
}

func (this_board Board) IsTherePawn(vertical int) bool {
	for i := range this_board.Cells[vertical] {
		cell := this_board.Cells[vertical][i]
		if cell != nil && cell.Type.Name == "Pawn" {
			return true
		}
	}
	return false
}

func (this_board Board) canPieceReachCell(vPieceCoord int, hPieceCoord int, vCellCoord int, hCellCoord int) bool {
	var isReached bool = true

	var piece = this_board.Cells[vPieceCoord][hPieceCoord]
	var moves = piece.Type.Moves
	var shiftSign = piece.getShiftSign()

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
			emptyCell := this_board.Cells[vMiddleCoord][hMiddleCoord] == nil
			if !emptyCell {
				isReached = false
				break
			}
		}
	}
	return isReached
}

func (this_board Board) GetInBoardFieldMoves(verticalCoord int, horizontalCoord int) (movesCoords [][2]int) {
	movesCoords = [][2]int{}

	var piece = this_board.Cells[verticalCoord][horizontalCoord]
	var shiftSign = piece.getShiftSign()
	var moves = piece.Type.Moves
	for _, move := range moves {
		var vMoveCoord = verticalCoord + move.HorizontalShift*shiftSign
		var hMoveCoord = horizontalCoord + move.VerticalShift*shiftSign

		var inBoardField bool = vMoveCoord >= 0 && vMoveCoord < len(this_board.Cells) && hMoveCoord >= 0 && hMoveCoord < len(this_board.Cells[vMoveCoord])
		if inBoardField {
			movesCoords = append(movesCoords, [2]int{vMoveCoord, hMoveCoord})
		}
	}
	return
}

func (this_board Board) GetPossibleMovesCoords(verticalCoord int, horizontalCoord int) (possibleMovesCoords [][2]int) {
	possibleMovesCoords = [][2]int{}

	var curPiece = this_board.Cells[verticalCoord][horizontalCoord]
	var inFieldMovesCoords = this_board.GetInBoardFieldMoves(verticalCoord, horizontalCoord)
	for _, coords := range inFieldMovesCoords {
		var emptyOrEnemyCell bool = this_board.Cells[coords[0]][coords[1]] == nil || this_board.Cells[coords[0]][coords[1]].Player != curPiece.Player
		if emptyOrEnemyCell {
			if this_board.canPieceReachCell(verticalCoord, horizontalCoord, coords[0], coords[1]) {
				possibleMovesCoords = append(possibleMovesCoords, [2]int{coords[0], coords[1]})
			}
		}
	}

	return
}

func (this_board Board) GetPossibleDropsCoords() (dropsCoords [][2]int) {
	dropsCoords = [][2]int{}

	for vDropCoord, verticalCells := range this_board.Cells {
		for hDropCoord, cell := range verticalCells {
			if cell == nil { // empty cell
				dropsCoords = append(dropsCoords, [2]int{vDropCoord, hDropCoord})
			}
		}
	}
	return
}

func (this_board Board) IsImportantPieceAttackedByPiece(attackerVerticalCoord int, attackerHorizontalCoord int) bool {
	movesCoords := this_board.GetPossibleMovesCoords(attackerVerticalCoord, attackerHorizontalCoord)
	for _, coords := range movesCoords {
		if this_board.Cells[coords[0]][coords[1]] != nil && this_board.Cells[coords[0]][coords[1]].Type.ImportantPiece {
			return true
		}
	}
	return false
}

func (this_board Board) GetImportantPieceCoordsForPlayer(player model.Player) [2]int {
	for v := range this_board.Cells {
		for h, piece := range this_board.Cells[v] {
			if piece != nil && piece.Type.ImportantPiece && piece.Player == player {
				return [2]int{v, h}
			}
		}
	}
	return [2]int{-100, -100} //should be impossible
}

func (this_board Board) GetCoordsOfAttackersOnCell(vCoordCell int, hCoordCell int) [][2]int {
	var attackers = [][2]int{}
	var attackedCell = this_board.Cells[vCoordCell][hCoordCell]

	for v := range this_board.Cells {
		for h, cell := range this_board.Cells[v] {
			if cell == nil { // empty cell
				continue
			}

			if cell.Player != attackedCell.Player {
				var movesCoords = this_board.GetPossibleMovesCoords(v, h)
				idx := slices.Index(movesCoords, [2]int{vCoordCell, hCoordCell})
				if idx != -1 {
					attackers = append(attackers, [2]int{v, h})
				}
			}
		}
	}
	return attackers
}

// TODO: как-то мерджнуть GetMovesCoordsOfAttackersOnCell и GetCoordsOfAttackersOnCell
func (this_board Board) GetMovesCoordsOfAttackersOnCell(vCoordCell int, hCoordCell int) [][2]int {
	var attackerMovesCoords = [][2]int{}
	var attackedCell = this_board.Cells[vCoordCell][hCoordCell]

	for v := range this_board.Cells {
		for h, attacker := range this_board.Cells[v] {
			if attacker == nil { // empty cell
				continue
			}

			var emptyOrEnemyAttackedCell = attackedCell == nil || attackedCell.Player != attacker.Player
			if emptyOrEnemyAttackedCell {
				var movesCoords = this_board.GetPossibleMovesCoords(v, h)
				idx := slices.Index(movesCoords, [2]int{vCoordCell, hCoordCell})
				if idx != -1 {
					attackerMovesCoords = append(attackerMovesCoords, movesCoords...)
				}
			}
		}
	}
	return attackerMovesCoords
}

// For test purposes
func (this_board Board) Print() {
	fmt.Println("-------------")
	for pieceType := range this_board.Inventories[model.Sente].pieces {
		piece := Piece{Type: *pieceType, Player: model.Sente}
		if piece.IsPromoted() {
			fmt.Print(string(piece.Type.Name[0]) + "+ ")
		} else {
			fmt.Print(string(piece.Type.Name[0]) + " ")
		}
	}
	fmt.Println()
	for _, verticalPieces := range this_board.Cells {
		for _, piece := range verticalPieces {
			if piece == nil {
				fmt.Print("-")
			} else if piece.IsPromoted() {
				fmt.Print(string(piece.Type.Name[0]) + "+")
			} else {
				fmt.Print(string(piece.Type.Name[0]))
			}
		}
		fmt.Println()
	}
	for pieceType := range this_board.Inventories[model.Gote].pieces {
		piece := Piece{Type: *pieceType, Player: model.Sente}
		if piece.IsPromoted() {
			fmt.Print(string(piece.Type.Name[0]) + "+ ")
		} else {
			fmt.Print(string(piece.Type.Name[0]) + " ")
		}
	}
	fmt.Println("\n-------------")
}
