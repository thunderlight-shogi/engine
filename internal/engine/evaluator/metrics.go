package evaluator

import (
	"math/rand"

	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
)

// Metrics weights
const MATERIAL_WEIGHT = 1.0
const ATTACK_COUNT_WEIGHT = 1.0
const PIECE_ADVANCEMENT_WEIGHT = 1.0
const DEFENDED_PIECES_WEIGHT = 1.0
const CHECK_WEIGHT = 1.0
const CHECKMATE_WEIGHT = 1.0
const KING_SAFETY_WEIGHT = 1.0

// Extra constants
const INVENTORY_MULTIPLIER = 1.5

func material(gameState *movegen.GameState) float32 {
	/*
		Sum of piece values of each side
		Pieces in inventory counted separately with its own weight
	*/
	var result float32 = rand.Float32()*10 - 5
	return result * MATERIAL_WEIGHT
}

func attackCount(gameState *movegen.GameState) float32 {
	/*
		Count of all cells that are attacked by player
	*/
	var result float32 = 0
	return result * ATTACK_COUNT_WEIGHT
}

func pieceAdvancement(gameState *movegen.GameState) float32 {
	/*
		How close pieces moved towards enemy's zone
	*/
	var result float32 = 0
	return result * PIECE_ADVANCEMENT_WEIGHT
}

func defendedPieces(gameState *movegen.GameState) float32 {
	/*
		Count of defended pieces
	*/
	var result float32 = 0
	return result * DEFENDED_PIECES_WEIGHT
}

func checkCheck(gameState *movegen.GameState) float32 {
	/*
		Is there check on the board
	*/
	var result float32 = 0
	return result * CHECK_WEIGHT
}

func checkCheckmate(gameState *movegen.GameState) float32 {
	/*
		Is there checkmate on the board
	*/
	var result float32 = 0
	return result * CHECKMATE_WEIGHT
}

func kingSafety(gameState *movegen.GameState) float32 {
	/*
		How safe is king:
			- how many friendly pieces around king
			- how many cells are defended around king
			- how many cells are attacked by enemy around king base
			- how many free cells can king go to
	*/
	var result float32 = 0
	return result * KING_SAFETY_WEIGHT
}
