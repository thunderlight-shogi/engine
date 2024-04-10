package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thunderlight-shogi/engine/internal/engine"
	"github.com/thunderlight-shogi/engine/internal/model"
)

func getPositionList() (pos_list []model.StartingPosition, err error) {
	pos_list = []model.StartingPosition{}
	db := model.GetDB()
	err = db.Find(&pos_list).Error
	return
}

func positionListHandler(w http.ResponseWriter, r *http.Request) {
	positions, err := getPositionList()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	enc.Encode(positions)
}

func startEngineHandler(w http.ResponseWriter, r *http.Request) {
	id_str := r.PathValue("id")

	var (
		id  uint64
		err error
	)

	id, err = strconv.ParseUint(id_str, 10, 0)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = engine.Start(uint(id))
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func Run() {
	http.HandleFunc("GET /position/list", positionListHandler)
	http.HandleFunc("POST /start/{id}", startEngineHandler)
	http.ListenAndServe(":80", nil)
}
