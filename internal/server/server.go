package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thunderlight-shogi/engine/internal/engine"
	"github.com/thunderlight-shogi/engine/internal/engine/board"
	"github.com/thunderlight-shogi/engine/internal/model"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, "text/html")
	w.Write([]byte("Hello, world!"))
}

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

	err = engine.Move(move)

	if err != nil {
		writeError(w, err)
		return
	}
}

func moveEngineHandler(w http.ResponseWriter, r *http.Request) {
	move, err := engine.GetMove()
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

}

func presetListHandler(w http.ResponseWriter, r *http.Request) {
	positions, err := getPositionList()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	setContentType(w, "application/json")
	enc := json.NewEncoder(w)

	enc.Encode(positions)
}
func presetGetHandler(w http.ResponseWriter, r *http.Request) {

}
func presetAddHandler(w http.ResponseWriter, r *http.Request) {

}
func presetUpdHandler(w http.ResponseWriter, r *http.Request) {

}
func presetDelHandler(w http.ResponseWriter, r *http.Request) {

}

func pieceListHandler(w http.ResponseWriter, r *http.Request) {

}
func pieceGetHandler(w http.ResponseWriter, r *http.Request) {

}
func pieceAddHandler(w http.ResponseWriter, r *http.Request) {

}
func pieceUpdHandler(w http.ResponseWriter, r *http.Request) {

}
func pieceDelHandler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	http.HandleFunc("GET /", indexHandler)             // Get main page
	http.HandleFunc("POST /start", startEngineHandler) // Start engine

	http.HandleFunc("POST /move/player", movePlayerHandler) // Player made a move
	http.HandleFunc("POST /move/engine", moveEngineHandler) // Get move from engine
	http.HandleFunc("POST /move/help", moveHelpHandler)     // Get a hint from engine

	http.HandleFunc("GET /preset/list", presetListHandler) // Get of all available presets
	http.HandleFunc("GET /preset/get", presetGetHandler)   // Get preset
	http.HandleFunc("POST /preset/add", presetAddHandler)  // Create new preset
	http.HandleFunc("POST /preset/upd", presetUpdHandler)  // Update existing preset
	http.HandleFunc("POST /preset/del", presetDelHandler)  // Delete preset

	http.HandleFunc("GET /piece/list", pieceListHandler) // Get list of all available pieces
	http.HandleFunc("GET /piece/get", pieceGetHandler)   // Get piece for preview
	http.HandleFunc("POST /piece/add", pieceAddHandler)  // Create a new piece
	http.HandleFunc("POST /piece/upd", pieceUpdHandler)  // Update existing piece
	http.HandleFunc("POST /piece/del", pieceDelHandler)  // Delete piece

	http.ListenAndServe(":80", nil)

	db := model.GetDB()
	_cool := model.PieceType{}
	db.Preload("Moves").Preload("PromotePiece").Preload("PromotePiece.Moves").First(&_cool, 13)

	str, _ := json.Marshal(_cool)
	fmt.Println(string(str))
	var _new model.PieceType
	json.Unmarshal(str, &_new)
	fmt.Println(_new)
}
