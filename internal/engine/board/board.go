package board

import (
	"fmt"
	"slices"

	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

type board struct {
	Cells               [9][9]*Piece // [File][Rank]
	Inventories         map[model.Player]Inventory
	reachableCellsCache [9][9][]Position
	possibleMovesCache  [9][9][]Position
}

type Board = *board

func Construct() (newby Board) {
	newby = new(board)

	newby.Inventories = make(map[model.Player]Inventory)
	newby.Inventories[model.Sente] = newInventory()
	newby.Inventories[model.Gote] = newInventory()

	return
}

func (this_board Board) Clone() (newby Board) {
	newby = new(board)

	for x := range newby.Cells {
		for y := range newby.Cells[x] {
			newby.Cells[x][y] = this_board.Cells[x][y]
		}
	}

	newby.Inventories = make(map[model.Player]Inventory)
	newby.Inventories[model.Sente] = this_board.Inventories[model.Sente].clone()
	newby.Inventories[model.Gote] = this_board.Inventories[model.Gote].clone()

	for x := range newby.reachableCellsCache {
		for y := range newby.reachableCellsCache[x] {
			newby.reachableCellsCache[x][y] = this_board.reachableCellsCache[x][y]
		}
	}

	for x := range newby.possibleMovesCache {
		for y := range newby.possibleMovesCache[x] {
			newby.possibleMovesCache[x][y] = this_board.possibleMovesCache[x][y]
		}
	}

	return
}

func (this_board Board) Equal(other Board) bool {
	for x := range this_board.Cells {
		for y := range this_board.Cells[x] {
			if this_board.Cells[x][y] == nil {
				if other.Cells[x][y] != nil {
					return false
				}
			} else if other.Cells[x][y] == nil {
				return false
			} else if *this_board.Cells[x][y] != *other.Cells[x][y] {
				return false
			}
		}
	}

	for key, value := range this_board.Inventories[model.Sente].pieces {
		if other.Inventories[model.Sente].pieces[key] != value {
			return false
		}
	}

	for key, value := range this_board.Inventories[model.Gote].pieces {
		if other.Inventories[model.Gote].pieces[key] != value {
			return false
		}
	}
	return true
}

func (this_board Board) At(pos Position) *Piece {
	return this_board.Cells[pos.File][pos.Rank]
}

func (this_board Board) Set(pos Position, piece *Piece) {
	this_board.Cells[pos.File][pos.Rank] = piece
}

func (this_board Board) MakeMove(fromPos, toPos Position, withPromotion bool) {
	cellMoveFrom := this_board.At(fromPos)
	cellMoveTo := this_board.At(toPos)

	var emptyOrFriendCell = cellMoveTo == nil || cellMoveFrom.Player == cellMoveTo.Player
	if !emptyOrFriendCell {
		this_board.Inventories[cellMoveFrom.Player].AddPiece(cellMoveTo)
	}

	this_board.Set(fromPos, nil)
	if withPromotion {
		this_board.Set(toPos, cellMoveFrom.GetPromotedPiece())
	} else {
		this_board.Set(toPos, cellMoveFrom)
	}

	// Updating cache
	this_board.updateCacheAfterMove(fromPos, toPos)
}

func (this_board Board) MakeDrop(pieceType *model.PieceType, player model.Player, dropPos Position) {
	this_board.Set(dropPos, this_board.Inventories[player].ExtractPieceToPlayer(pieceType, player))

	// Updating cache
	this_board.updateCacheAfterDrop(dropPos)
}

func (this_board Board) IsTherePawn(file int) bool {
	for i := range this_board.Cells[file] {
		cell := this_board.Cells[file][i]
		if cell != nil && cell.Type.Name == "Pawn" {
			return true
		}
	}
	return false
}

func (this_board Board) canPieceReachCell(piecePos, cellPos Position) bool {
	var isReached bool = true

	pieceFile, pieceRank := piecePos.Get()
	cellFile, cellRank := cellPos.Get()

	var piece = this_board.Cells[pieceFile][pieceRank]
	var moves = piece.Type.Moves
	var shiftSign = piece.getShiftSign()

	origin := graphics.NewPoint(pieceFile, pieceRank)
	end := graphics.NewPoint(cellFile, cellRank)
	coordsBetween := graphics.GetLinePoints(origin, end)
	for i := 1; i < len(coordsBetween)-1; i++ { // skipping origin and end
		vMiddleCoord, hMiddleCoord := coordsBetween[i].Coordinates()
		hShift, vShift := vMiddleCoord-pieceFile, hMiddleCoord-pieceRank
		// maybe better to replace with simple loop
		idx := slices.IndexFunc(moves, func(move model.Move) bool {
			return move.FileShift*shiftSign == hShift && move.RankShift*shiftSign == vShift
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

func (this_board Board) GetPieceMovesToBoardField(piecePos Position) (movesPositions []Position) {
	//fmt.Printf("piecePos: %v\n", piecePos)
	movesPositions = []Position{}
	pieceFile, pieceRank := piecePos.Get()

	var piece = this_board.Cells[pieceFile][pieceRank]
	//fmt.Printf("piece: %v\n", piece.Type)
	var shiftSign = piece.getShiftSign()
	var moves = piece.Type.Moves
	for _, move := range moves {
		var moveFile = pieceFile + move.FileShift*shiftSign
		var moveRank = pieceRank + move.RankShift*shiftSign

		var inBoardField bool = moveFile >= 0 && moveFile < len(this_board.Cells) && moveRank >= 0 && moveRank < len(this_board.Cells[moveFile])
		if inBoardField {
			movesPositions = append(movesPositions, NewPos(moveFile, moveRank))
		}
	}
	return
}

func (this_board Board) GetPieceReachableCells(piecePos Position, useCache bool) (reachableCellsPositions []Position) {
	if useCache {
		return this_board.reachableCellsCache[piecePos.File][piecePos.Rank]
	}

	reachableCellsPositions = []Position{}

	var inFieldMovesPositions = this_board.GetPieceMovesToBoardField(piecePos)
	for _, movePos := range inFieldMovesPositions {
		if this_board.canPieceReachCell(piecePos, movePos) {
			reachableCellsPositions = append(reachableCellsPositions, movePos)
		}
	}
	return
}

func (this_board Board) GetPiecePossibleMoves(piecePos Position, useCache bool) (possibleMovesPositions []Position) {
	if useCache {
		return this_board.possibleMovesCache[piecePos.File][piecePos.Rank]
	}

	possibleMovesPositions = []Position{}
	pieceFile, pieceRank := piecePos.Get()

	var curPiece = this_board.Cells[pieceFile][pieceRank]
	var reachableCellsPositions = this_board.GetPieceReachableCells(piecePos, true)
	for _, movePos := range reachableCellsPositions {
		var moveCell = this_board.Cells[movePos.File][movePos.Rank]
		var emptyOrEnemyCell bool = moveCell == nil || moveCell.Player != curPiece.Player
		if emptyOrEnemyCell {
			possibleMovesPositions = append(possibleMovesPositions, movePos)
		}
	}
	return
}

func (this_board Board) GetPossibleDropsCoords() (dropsCoords []Position) {
	dropsCoords = []Position{}

	this_board.IterateEmptyCells(func(pos Position) {
		dropsCoords = append(dropsCoords, pos)
	})
	return
}

func (this_board Board) GetKingPosition(kingPlayer model.Player) Position {
	for x := range this_board.Cells {
		for y, piece := range this_board.Cells[x] {
			if piece != nil && piece.Type.ImportantPiece && piece.Player == kingPlayer {
				return NewPos(x, y)
			}
		}
	}
	this_board.Print()
	panic("The king is not on the board")
}

func (this_board Board) IsCellAttacked(attackerPlayer model.Player, cellPos Position) bool {
	var isCellAttacked bool = false
	this_board.IterateBoardPiecesWithEarlyExit(attackerPlayer, func(piece *Piece, pos Position) bool {
		var movesPositions = this_board.GetPieceReachableCells(pos, true)
		idx := slices.Index(movesPositions, cellPos)
		if idx != -1 {
			isCellAttacked = true
			return true
		}
		return false
	})

	return isCellAttacked
}

func (this_board Board) IsKingAttacked(kingPlayer model.Player) bool {
	var kingPos = this_board.GetKingPosition(kingPlayer)
	var attackerPlayer = this_board.At(kingPos).GetAttackerPlayer()
	result := this_board.IsCellAttacked(attackerPlayer, kingPos)
	return result
}

func (this_board Board) IsKingAttackedByPiece(kingPlayer model.Player, attackerPos Position) bool {
	attackerMovesPositions := this_board.GetPiecePossibleMoves(attackerPos, true)
	for _, movePos := range attackerMovesPositions {
		moveFile, moveRank := movePos.Get()
		var piece = this_board.Cells[moveFile][moveRank]
		var isImportantPiece bool = piece != nil && piece.Type.ImportantPiece
		if isImportantPiece && piece.Player == kingPlayer {
			return true
		}
	}
	return false
}

func (this_board Board) GetPositionsOfAttackersOnCell(attackerPlayer model.Player, cellPos Position) []Position {
	var attackers = []Position{}

	this_board.IterateBoardPieces(attackerPlayer, func(piece *Piece, pos Position) {
		var movesPositions = this_board.GetPieceReachableCells(pos, true)
		idx := slices.Index(movesPositions, cellPos)
		if idx != -1 {
			attackers = append(attackers, pos)
		}
	})
	return attackers
}

// TODO: Объединить в GetPiecePossibleMoves
func (this_board Board) GetKingPossibleMoves(kingPos Position) []Position {
	var kingMovesPositions = []Position{}
	var king = this_board.At(kingPos)
	var attackerPlayer = king.GetAttackerPlayer()

	var potentialMovesPositions = this_board.GetPiecePossibleMoves(kingPos, false)
	for _, pos := range potentialMovesPositions {
		if !this_board.IsCellAttacked(attackerPlayer, pos) {
			kingMovesPositions = append(kingMovesPositions, pos)
		}
	}
	return kingMovesPositions
}

// For test purposes
func (this_board Board) Print() {
	fmt.Println("----------------------")
	for pieceType, count := range this_board.Inventories[model.Gote].pieces {
		for range count {
			fmt.Print(string(pieceType.Kanji))
		}
	}
	fmt.Println()

	for i := range 9 {
		fmt.Print(" ", 8-i, " ")
	}
	fmt.Println()

	for rank := range 9 {
		for file := range 9 {
			var piece = this_board.Cells[8-file][rank]
			if piece == nil {
				fmt.Print(" - ")
			} else {
				str := ""
				switch piece.Player {
				case model.Sente:
					str += "☖"
				case model.Gote:
					str += "☗"
				}
				if piece.IsPromoted() {
					str += "\x1b[96m" + string(piece.Type.Kanji) + "\x1b[97m"
				} else {
					str += string(piece.Type.Kanji)
				}

				fmt.Print(str)
			}
		}
		fmt.Print(" ", rank, "\n")
	}

	for pieceType, count := range this_board.Inventories[model.Sente].pieces {
		for range count {
			fmt.Print(string(pieceType.Kanji))
		}
	}
	fmt.Println("\n----------------------")
}
