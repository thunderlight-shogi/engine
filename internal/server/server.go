package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/thunderlight-shogi/engine/internal/engine"
)

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	setContentType(w, "text/html")
	w.Write([]byte("Hello, world!"))
}

func Run() {
	http.HandleFunc("GET /", indexHandler)             // Get main page
	http.HandleFunc("POST /start", startEngineHandler) // Start engine

	http.HandleFunc("POST /move/player", indexHandler) // Player made a move
	http.HandleFunc("POST /move/engine", indexHandler) // Get move from engine
	http.HandleFunc("POST /move/help", indexHandler)   // Get a hint from engine

	http.HandleFunc("GET /preset/list", presetListHandler) // Get of all available presets
	http.HandleFunc("GET /preset/get", indexHandler)       // Get preset
	http.HandleFunc("POST /preset/add", indexHandler)      // Create new preset
	http.HandleFunc("POST /preset/upd", indexHandler)      // Update existing preset
	http.HandleFunc("POST /preset/del", indexHandler)      // Delete preset

	http.HandleFunc("POST /piece/list", indexHandler) // Get list of all available pieces
	http.HandleFunc("POST /piece/get", indexHandler)  // Get piece for preview
	http.HandleFunc("POST /piece/add", indexHandler)  // Create a new piece
	http.HandleFunc("POST /piece/upd", indexHandler)  // Update existing piece
	http.HandleFunc("POST /piece/del", indexHandler)  // Delete piece

	http.ListenAndServe(":80", nil)
}
