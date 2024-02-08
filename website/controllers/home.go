package website

import (
	"io"
	"net/http"
)

func Home(response http.ResponseWriter, request *http.Request, parameters map[string]string) {
	io.WriteString(response, "<h1>It works!</h1>")
}
