package server

import (
	"net/http"

	"github.com/thunderlight-shogi/engine/internal/model"
)

func setContentType(w http.ResponseWriter, con_type string) {
	w.Header().Add("Content-Type", con_type)
}

func getPositionList() (pos_list []model.Preset, err error) {
	pos_list = []model.Preset{}
	db := model.GetDB()
	err = db.Find(&pos_list).Error
	return
}
