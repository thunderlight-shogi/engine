package evaluator

import (
	"fmt"
	"reflect"
	"runtime"

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

func evaluate(gameState *movegen.GameState) float32 {
	/*
		Does evaluation of game state by summing up each metric
	*/
	var result float32 = 0
	for _, metric := range used_metrics {
		result += metric(gameState, model.Sente)
		result -= metric(gameState, model.Sente)
	}
	return result
}

func evaluation_report(gameState *movegen.GameState) string {
	/*
		Returns info about each metric for debug purposes
	*/
	report := "Sente:\n"
	for _, metric := range used_metrics {
		func_name := runtime.FuncForPC(reflect.ValueOf(metric).Pointer()).Name()
		metric_value := metric(gameState, model.Sente)
		report += fmt.Sprintf("%s: %.2f", func_name, metric_value)
	}
	report += "\nGote:\n"
	for _, metric := range used_metrics {
		func_name := runtime.FuncForPC(reflect.ValueOf(metric).Pointer()).Name()
		metric_value := metric(gameState, model.Gote)
		report += fmt.Sprintf("%s: %.2f", func_name, metric_value)
	}
	return report
}
