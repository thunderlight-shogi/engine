package evaluator

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/thunderlight-shogi/engine/internal/engine/movegen"
)

// Metrics used when evaluating board
var used_metrics = [...]func(*movegen.GameState) float32{
	material,
	attackCount,
	pieceAdvancement,
	defendedPieces,
	checkCheck,
	checkCheckmate,
	kingSafety,
}

func evaluate(gameState *movegen.GameState) float32 {
	var result float32 = 0
	for _, metric := range used_metrics {
		result += metric(gameState)
	}
	return result
}

func evaluation_report(gameState *movegen.GameState) string {
	report := ""
	for _, metric := range used_metrics {
		func_name := runtime.FuncForPC(reflect.ValueOf(metric).Pointer()).Name()
		metric_value := metric(gameState)
		report += fmt.Sprintf("%s: %.2f", func_name, metric_value)
	}
	return report
}
