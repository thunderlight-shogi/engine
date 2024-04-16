package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/thunderlight-shogi/engine/internal/engine"
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

func startEngineHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		writeError(w, err)
		return
	}

	var preset model.Preset
	err = json.Unmarshal(body, &preset)

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
	move := engine.GetHelp()

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

	err := db.Find(&pieces).Error
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
	http.HandleFunc("POST /start", startEngineHandler) // Start engine

	http.HandleFunc("POST /move/player", movePlayerHandler) // Player made a move
	http.HandleFunc("POST /move/engine", moveEngineHandler) // Get move from engine
	http.HandleFunc("POST /move/help", moveHelpHandler)     // Get a hint from engine

	http.HandleFunc("GET /preset/list", presetListHandler) // Get of all available presets
	http.HandleFunc("POST /preset/get", presetGetHandler)  // Get preset
	http.HandleFunc("POST /preset/add", presetAddHandler)  // Create new preset
	http.HandleFunc("POST /preset/upd", presetUpdHandler)  // Update existing preset
	http.HandleFunc("POST /preset/del", presetDelHandler)  // Delete preset

	http.HandleFunc("GET /piece/list", pieceListHandler) // Get list of all available pieces
	http.HandleFunc("POST /piece/get", pieceGetHandler)  // Get piece for preview
	http.HandleFunc("POST /piece/add", pieceAddHandler)  // Create a new piece
	http.HandleFunc("POST /piece/upd", pieceUpdHandler)  // Update existing piece
	http.HandleFunc("POST /piece/del", pieceDelHandler)  // Delete piece

	http.HandleFunc("GET /move_type/get", moveTypeGetHandler) // Get available piece types

	http.ListenAndServe(":88", nil)
}
