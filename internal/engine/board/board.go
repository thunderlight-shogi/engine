package board

import (
	"fmt"
	"slices"

	"github.com/thunderlight-shogi/engine/internal/model"
	"github.com/thunderlight-shogi/engine/pkg/graphics"
)

type Position struct {
	File int `json:"file"`
	Rank int `json:"rank"`
}

type MoveType uint

const (
	Moving MoveType = iota
	Attacking
	Dropping
	PromotionMoving
	PromotionAttacking
)

type Move struct {
	OldCoords Position         `json:"old_pos"`
	NewCoords Position         `json:"new_pos"`
	PieceType *model.PieceType `json:"piece_type"`
	MoveType  MoveType         `json:"move_type"`
}

func NewPos(file, rank int) Position {
	return Position{File: file, Rank: rank}
}

func (pos Position) Get() (int, int) {
	return pos.File, pos.Rank
}

func (pos Position) GetFile() int {
	return pos.File
}

func (pos Position) GetRank() int {
	return pos.Rank
}

type Piece struct {
	Type   *model.PieceType
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

func (this_piece *Piece) GetPromotedPiece() *Piece {
	return &Piece{Type: this_piece.Type.PromotePiece, Player: this_piece.Player}
}

func (this_piece *Piece) GetAttackerPlayer() (attackerPlayer model.Player) {
	if this_piece.Player == model.Sente {
		attackerPlayer = model.Gote
	} else {
		attackerPlayer = model.Sente
	}
	return
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

func (this_board Board) At(pos Position) *Piece {
	return this_board.Cells[pos.File][pos.Rank]
}

func (this_board Board) Set(pos Position, piece *Piece) {
	this_board.Cells[pos.File][pos.Rank] = piece
}

func (this_board Board) GetPromotionZone(player model.Player) []int {
	if player == model.Sente {
		return PromotionZoneForSente
	} else {
		return PromotionZoneForGote
	}
}

func (this_board Board) IterateInventory(
	player model.Player,
	callback func(piece *model.PieceType),
) {
	/*
		Iterates over player's inventory in current game state
		and for each figure calls callback function
	*/
	playerInventory := this_board.Inventories[player]
	for _, pieceType := range playerInventory.Pieces() {
		count := playerInventory.CountPiece(pieceType)
		for range count {
			callback(pieceType)
		}
	}
}

func (this_board Board) IterateBoardPieces(
	player model.Player,
	callback func(piece *Piece, pos Position),
) {
	/*
		Iterates over player's pieces on board in current game state
		and for each figure calls callback function
	*/
	for x, column := range this_board.Cells {
		for y, cell := range column {
			if cell != nil && cell.Player == player { // If player's cell
				callback(cell, NewPos(x, y))
			}
		}
	}
}

func (this_board Board) IterateEmptyCells(
	callback func(pos Position),
) {
	/*
		Iterates over all empty cells on board
		and for each calls callback function
	*/
	for x, column := range this_board.Cells {
		for y, cell := range column {
			if cell == nil {
				callback(NewPos(x, y))
			}
		}
	}
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

func (this_board Board) GetPieceReachableMoves(piecePos Position) (possibleMovesPositions []Position) {
	possibleMovesPositions = []Position{}

	var inFieldMovesPositions = this_board.GetPieceMovesToBoardField(piecePos)
	for _, movePos := range inFieldMovesPositions {
		if this_board.canPieceReachCell(piecePos, movePos) {
			possibleMovesPositions = append(possibleMovesPositions, movePos)
		}
	}
	return
}

func (this_board Board) GetPiecePossibleMoves(piecePos Position) (possibleMovesPositions []Position) {
	possibleMovesPositions = []Position{}
	pieceFile, pieceRank := piecePos.Get()

	var curPiece = this_board.Cells[pieceFile][pieceRank]
	var reachableMovesPositions = this_board.GetPieceReachableMoves(piecePos)
	for _, movePos := range reachableMovesPositions {
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

func (this_board Board) IsKingAttackedByPiece(attackerPos Position) bool {
	attackerMovesPositions := this_board.GetPiecePossibleMoves(attackerPos)
	for _, movePos := range attackerMovesPositions {
		moveFile, moveRank := movePos.Get()
		var isImportantPiece bool = this_board.Cells[moveFile][moveRank] != nil && this_board.Cells[moveFile][moveRank].Type.ImportantPiece
		if isImportantPiece {
			return true
		}
	}
	return false
}

func (this_board Board) GetKingPositionForPlayer(player model.Player) Position {
	for x := range this_board.Cells {
		for y, piece := range this_board.Cells[x] {
			if piece != nil && piece.Type.ImportantPiece && piece.Player == player {
				return NewPos(x, y)
			}
		}
	}
	panic("The king is not on the board")
}

func (this_board Board) GetPositionsOfAttackersOnCell(attackerPlayer model.Player, cellPos Position) []Position {
	var attackers = []Position{}

	this_board.IterateBoardPieces(attackerPlayer, func(piece *Piece, pos Position) {
		var movesPositions = this_board.GetPieceReachableMoves(pos)
		idx := slices.Index(movesPositions, cellPos)
		if idx != -1 {
			attackers = append(attackers, pos)
		}
	})
	return attackers
}

func (this_board Board) GetKingPossibleMoves(kingPos Position) []Position {
	var kingMovesPositions = []Position{}
	var king = this_board.Cells[kingPos.GetFile()][kingPos.GetRank()]
	var attackerPlayer = king.GetAttackerPlayer()

	var potentialMovesPositions = this_board.GetPiecePossibleMoves(kingPos)
	for _, pos := range potentialMovesPositions {
		if len(this_board.GetPositionsOfAttackersOnCell(attackerPlayer, pos)) == 0 {
			kingMovesPositions = append(kingMovesPositions, pos)
		}
	}
	return kingMovesPositions
}

// For test purposes
func (this_board Board) Print() {
	fmt.Println("-------------")
	for pieceType, count := range this_board.Inventories[model.Gote].pieces {
		piece := Piece{Type: pieceType, Player: model.Sente}
		for i := uint(0); i < count; i++ {
			if piece.IsPromoted() {
				fmt.Print(string(piece.Type.Name[0]) + "+ ")
			} else {
				fmt.Print(string(piece.Type.Name[0]) + " ")
			}
		}
	}
	fmt.Println()

	for i := range 9 {
		for j := range 9 {
			var piece = this_board.Cells[8-j][i]
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
				str += string(piece.Type.Kanji)

				fmt.Print(str)
			}
		}
		fmt.Println()
	}

	for pieceType, count := range this_board.Inventories[model.Sente].pieces {
		piece := Piece{Type: pieceType, Player: model.Sente}
		for i := uint(0); i < count; i++ {
			if piece.IsPromoted() {
				fmt.Print(string(piece.Type.Name[0]) + "+ ")
			} else {
				fmt.Print(string(piece.Type.Name[0]) + " ")
			}
		}
	}
	fmt.Println("\n-------------")
}
