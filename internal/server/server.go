package server

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"

	"github.com/gorilla/handlers"

	"github.com/thunderlight-shogi/engine/internal/engine"
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

func startEngineHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(string(body));

	if err != nil {
		writeError(w, err)
		return
	}

	var preset model.Preset
	err = json.Unmarshal(body, &preset)

	fmt.Println(preset.Id)

	if err != nil {
		writeError(w, err)
		return
	}

	if preset.Id > 0 {
		engine.Start(preset.Id)
	} else {
		writeError(w, err)
		return
	}
}

func movePlayerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		writeError(w, err)
		return
	}

	var move board.Move
	json.Unmarshal(body, &move)

	move.PieceType, err = engine.FindPiece(move.PieceType.Id)
	if err != nil {
		writeError(w, err)
		return
	}

	err = engine.Move(move)

	if err != nil {
		writeError(w, err)
		return
	}
}

func moveEngineHandler(w http.ResponseWriter, r *http.Request) {
	move, err := engine.EngineMove()
	if err != nil {
		writeError(w, err)
		return
	}

	body, err := json.Marshal(move)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(body)
}

func moveHelpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("move/help is requested. Waiting for the engine to response...")

	move := engine.GetHelp()

	fmt.Println("move/help is finished.")

	body, err := json.Marshal(move)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(body)
}

func presetListHandler(w http.ResponseWriter, r *http.Request) {
	presets := make([]model.Preset, 0)
	db := model.GetDB()

	err := db.Find(&presets).Error
	if err != nil {
		writeError(w, err)
		return
	}

	str, err := json.Marshal(presets)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(str)
}
func presetGetHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var preset model.Preset

	err = json.Unmarshal(body, &preset)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Preload("Pieces").Preload("Pieces.PieceType").First(&preset).Error
	if err != nil {
		writeError(w, err)
		return
	}

	str, err := json.Marshal(preset)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(str)
}
func presetAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var preset model.Preset

	err = json.Unmarshal(body, &preset)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Create(preset).Error
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}
func presetUpdHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var preset model.Preset

	err = json.Unmarshal(body, &preset)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Save(preset).Error
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}
func presetDelHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var preset model.Preset

	err = json.Unmarshal(body, &preset)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Delete(preset).Error
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}

func pieceListHandler(w http.ResponseWriter, r *http.Request) {
	pieces := make([]model.PieceType, 0)
	db := model.GetDB()

	err := db.
		Preload("PromotePiece").
		Find(&pieces).Error

	if err != nil {
		writeError(w, err)
		return
	}

	str, err := json.Marshal(pieces)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(str)
}
func pieceGetHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var piece model.PieceType

	err = json.Unmarshal(body, &piece)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Preload("Moves").Preload("PromotePiece").Preload("PromotePiece.Moves").First(&piece).Error
	if err != nil {
		writeError(w, err)
		return
	}

	str, err := json.Marshal(piece)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(str)
}
func pieceAddHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var piece model.PieceType

	err = json.Unmarshal(body, &piece)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Create(piece).Error
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}
func pieceUpdHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var piece model.PieceType

	err = json.Unmarshal(body, &piece)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Save(piece).Error
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}
func pieceDelHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		writeError(w, err)
		return
	}

	db := model.GetDB()
	var piece model.PieceType

	err = json.Unmarshal(body, &piece)
	if err != nil {
		writeError(w, err)
		return
	}

	err = db.Delete(piece).Error
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(200)
}

func moveTypeGetHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]board.MoveType)

	m["moving"] = board.Moving
	m["attacking"] = board.Attacking
	m["dropping"] = board.Dropping
	m["promotion_moving"] = board.PromotionMoving
	m["promotion_attacking"] = board.PromotionAttacking

	str, err := json.Marshal(m)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Write(str)
}

func Run() {
	fmt.Println("[!] The Thunderlight RestAPI Server is about to start.")

	http.Handle("/start", handlers.CORS()(http.HandlerFunc(startEngineHandler)))
    http.Handle("/move/player", handlers.CORS()(http.HandlerFunc(movePlayerHandler)))
    http.Handle("/move/engine", handlers.CORS()(http.HandlerFunc(moveEngineHandler)))
    http.Handle("/move/help", handlers.CORS()(http.HandlerFunc(moveHelpHandler)))
    http.Handle("/preset/list", handlers.CORS()(http.HandlerFunc(presetListHandler)))
    http.Handle("/preset/get", handlers.CORS()(http.HandlerFunc(presetGetHandler)))
    http.Handle("/preset/add", handlers.CORS()(http.HandlerFunc(presetAddHandler)))
    http.Handle("/preset/upd", handlers.CORS()(http.HandlerFunc(presetUpdHandler)))
    http.Handle("/preset/del", handlers.CORS()(http.HandlerFunc(presetDelHandler)))
    http.Handle("/piece/list", handlers.CORS()(http.HandlerFunc(pieceListHandler)))
    http.Handle("/piece/get", handlers.CORS()(http.HandlerFunc(pieceGetHandler)))
    http.Handle("/piece/add", handlers.CORS()(http.HandlerFunc(pieceAddHandler)))
    http.Handle("/piece/upd", handlers.CORS()(http.HandlerFunc(pieceUpdHandler)))
    http.Handle("/piece/del", handlers.CORS()(http.HandlerFunc(pieceDelHandler)))
    http.Handle("/move_type/get", handlers.CORS()(http.HandlerFunc(moveTypeGetHandler)))
	
	http.ListenAndServe(":88", nil)
}
