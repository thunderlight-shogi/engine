package evaluator

import (
	"fmt"

	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
	"github.com/thunderlight-shogi/engine/internal/model"
)

// Metrics used when evaluating board
var used_metrics = [...]func(*movegen.GameState, model.Player) float32{
	material,
	attackCount,
	pieceAdvancement,
	defendedPieces,
	checkCheck,
	checkCheckmate,
	kingGuardsCount,
	kingDefenceRadius1,
	kingDefenceRadius2,
	kingAttackRadius1,
	kingAttackRadius2,
	kingFreeCells,
}

var titles = [...]string{
	"material",
	"attackCount",
	"pieceAdvancement",
	"defendedPieces",
	"checkCheck",
	"checkCheckmate",
	"kingGuardsCount",
	"kingDefenceRadius1",
	"kingDefenceRadius2",
	"kingAttackRadius1",
	"kingAttackRadius2",
	"kingFreeCells",
}

func Evaluate(gameState *movegen.GameState) float32 {
	/*
		Does evaluation of game state by summing up each metric
	*/
	var result float32 = 0
	for _, metric := range used_metrics {
		result += metric(gameState, model.Sente)
		result -= metric(gameState, model.Gote)
	}
	return result
}

func Evaluation_report(gameState *movegen.GameState) string {
	/*
		Returns info about each metric for debug purposes
	*/
	report := "Sente:\n"
	for i, metric := range used_metrics {
		func_name := titles[i]
		metric_value := metric(gameState, model.Sente)
		report += fmt.Sprintf("%s: %.2f\n", func_name, metric_value)
	}
	report += "\nGote:\n"
	for i, metric := range used_metrics {
		func_name := titles[i]
		metric_value := metric(gameState, model.Gote)
		report += fmt.Sprintf("%s: %.2f\n", func_name, metric_value)
	}
	return report
}
