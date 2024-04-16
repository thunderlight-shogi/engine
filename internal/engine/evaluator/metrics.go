package evaluator

import (
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
	"github.com/thunderlight-shogi/engine/internal/model"
)

// Metrics weights
const MATERIAL_WEIGHT = 1.0
const ATTACK_COUNT_WEIGHT = 1.0
const PIECE_ADVANCEMENT_WEIGHT = 1.0
const DEFENDED_PIECES_WEIGHT = 1.0
const CHECK_WEIGHT = 10.0
const CHECKMATE_WEIGHT = 99999.0

const KING_GUARDS_COUNT_WEIGHT = 1.0
const KING_DEFENCE_RADIUS1_WEIGHT = 2.0
const KING_DEFENCE_RADIUS2_WEIGHT = 1.0
const KING_ATTACK_RADIUS1_WEIGHT = -2.0
const KING_ATTACK_RADIUS2_WEIGHT = -1.0
const KING_FREE_CELLS_WEIGHT = 3.0

// Extra constants
const INVENTORY_MULTIPLIER = 1.5
const MAX_KING_GUARDS_DISTANCE = 2

func material(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		Sum of piece values
		Pieces in inventory counted separately with its own weight
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	curBoard.IterateInventory(player, func(piece *model.PieceType) {
		result += float32(piece.Cost) * INVENTORY_MULTIPLIER
	})
	curBoard.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		result += float32(piece.Type.Cost)
	})

	return result * MATERIAL_WEIGHT
}

func attackCount(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		Count of all cells that are attacked by player
	*/
	attackMatrix := createAttackMatrix(gameState.Board, player)
	var result float32 = float32(sumOf2DMatrix(attackMatrix))

	return result * ATTACK_COUNT_WEIGHT
}

func pieceAdvancement(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How close pieces moved towards enemy's zone
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	curBoard.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		var y = pos.GetRank()
		// Calculating y-offset from player's base
		var baseYOffset int
		if player == model.Sente { // First player. Base starts at 8 and goes to 0
			baseYOffset = 8 - y
		} else { // Second player. Base starts at 0 and goes to 8
			baseYOffset = y
		}
		advancementScore := min(baseYOffset, 6) // Assume that last 3 rows have same value
		advancementScore *= int(piece.Type.Cost / 3)
		result += float32(advancementScore)
	})

	return result * PIECE_ADVANCEMENT_WEIGHT
}

func defendedPieces(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		Count of defended pieces
	*/
	defendMatrix := createDefendMatrix(gameState.Board, player)
	var result float32 = float32(sumOf2DMatrix(defendMatrix))

	return result * DEFENDED_PIECES_WEIGHT
}

func checkCheck(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		Has player checked his opponent?
	*/
	isOpponentChecked := gameState.KingUnderAttack && gameState.CurMovePlayer != player
	var result float32
	if isOpponentChecked {
		result = 1
	} else {
		result = 0
	}
	return result * CHECK_WEIGHT
}

func checkCheckmate(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		Has player checkmated his opponent?
	*/

	// TODO: Вызывать GeneratePossibleStates довольно затратно только для того, чтобы проверить мат. Оптимизировать
	// 		 (можно проверять, находится ли король под шахом и только тогда генерить возможные геймстейты)
	isOpponentCheckmated := len(gameState.GeneratePossibleStates()) == 0 && gameState.CurMovePlayer != player
	var result float32
	if isOpponentCheckmated {
		result = 1
	} else {
		result = 0
	}
	return result * CHECKMATE_WEIGHT
}

// ====================== KING SAFETY METRICS ======================

// TODO: Как-то объединить метрики. У них почти идентичный код

func kingGuardsCount(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How many friendly pieces around king
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	kingCoords := curBoard.GetKingPositionForPlayer(player)
	kingX, kingY := kingCoords.Get()

	curBoard.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		x, y := pos.Get()
		distance := chebyshevDistance(kingX, kingY, x, y)
		if distance <= MAX_KING_GUARDS_DISTANCE {
			result += 1
		}
	})

	return result * KING_GUARDS_COUNT_WEIGHT
}

func kingDefenceRadius1(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How much friendly pieces are defended around king in radius 1
		Counting all defended cells around king
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	defendMatrix := createDefendMatrix(curBoard, player)
	kingCoords := gameState.Board.GetKingPositionForPlayer(player)
	kingX, kingY := kingCoords.Get()

	curBoard.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		x, y := pos.Get()
		distance := chebyshevDistance(kingX, kingY, x, y)
		if distance == 1 {
			result += float32(defendMatrix[x][y])
		}
	})

	return result * KING_DEFENCE_RADIUS1_WEIGHT
}

func kingDefenceRadius2(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How many friendly pieces are defended around king in radius 2
		Counting all defended cells around king
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	defendMatrix := createDefendMatrix(curBoard, player)
	kingCoords := gameState.Board.GetKingPositionForPlayer(player)
	kingX, kingY := kingCoords.Get()

	curBoard.IterateBoardPieces(player, func(piece *board.Piece, pos board.Position) {
		x, y := pos.Get()
		distance := chebyshevDistance(kingX, kingY, x, y)
		if distance == 2 {
			result += float32(defendMatrix[x][y])
		}
	})

	return result * KING_DEFENCE_RADIUS2_WEIGHT
}

func kingAttackRadius1(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How many cells are attacked by enemy around king in radius 1
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	attackMatrix := createAttackMatrix(curBoard, player)
	kingCoords := gameState.Board.GetKingPositionForPlayer(player)
	kingX, kingY := kingCoords.Get()

	curBoard.IterateEmptyCells(func(pos board.Position) {
		x, y := pos.Get()
		distance := chebyshevDistance(kingX, kingY, x, y)
		if distance == 1 {
			result += float32(attackMatrix[x][y])
		}
	})

	return result * KING_ATTACK_RADIUS1_WEIGHT
}

func kingAttackRadius2(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How many cells are attacked by enemy around king in radius 2
	*/
	var curBoard = gameState.Board
	var result float32 = 0

	attackMatrix := createAttackMatrix(curBoard, player)
	kingCoords := gameState.Board.GetKingPositionForPlayer(player)
	kingX, kingY := kingCoords.Get()

	curBoard.IterateEmptyCells(func(pos board.Position) {
		x, y := pos.Get()
		distance := chebyshevDistance(kingX, kingY, x, y)
		if distance == 2 {
			result += float32(attackMatrix[x][y])
		}
	})

	return result * KING_ATTACK_RADIUS2_WEIGHT
}

func kingFreeCells(
	gameState *movegen.GameState,
	player model.Player,
) float32 {
	/*
		How many free cells can king go to
	*/
	kingPos := gameState.Board.GetKingPositionForPlayer(player)

	var result float32 = float32(len(gameState.Board.GetKingPossibleMoves(kingPos)))

	return result * KING_FREE_CELLS_WEIGHT
}
