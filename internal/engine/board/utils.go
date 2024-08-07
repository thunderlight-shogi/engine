package board

import "github.com/thunderlight-shogi/engine/internal/model"

var PromotionZoneForSente = []int{0, 1, 2}
var PromotionZoneForGote = []int{6, 7, 8}

func GetPromotionZone(player model.Player) []int {
	if player == model.Sente {
		return PromotionZoneForSente
	} else {
		return PromotionZoneForGote
	}
}
