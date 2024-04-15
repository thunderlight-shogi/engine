package server

import (
	"net/http"
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
