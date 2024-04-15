package server

import (
	"net/http"

	"github.com/thunderlight-shogi/engine/internal/model"
)

func setContentType(w http.ResponseWriter, con_type string) {
	w.Header().Add("Content-Type", con_type)
}

func jsonType(w http.ResponseWriter) {
	setContentType(w, "application/json")
}

func plainType(w http.ResponseWriter) {
	setContentType(w, "text/plain")
}

func htmlType(w http.ResponseWriter) {
	setContentType(w, "text/html")
}

func writeError(w http.ResponseWriter, err error) {
	plainType(w)
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

func getPositionList() (pos_list []model.Preset, err error) {
	pos_list = []model.Preset{}
	db := model.GetDB()
	err = db.Find(&pos_list).Error
	return
}
