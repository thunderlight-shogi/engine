package website

import (
	"fmt"
	"io"
	"net/http"
)

func Editor(response http.ResponseWriter, request *http.Request, parameters map[string]string) {
	id, ok := parameters["id"]

	if !ok {
		io.WriteString(response, "ID should be set.")
	}

	io.WriteString(response, fmt.Sprintf("ID is %s", id))
}
