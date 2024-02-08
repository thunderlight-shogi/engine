package website

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const port int = 9999

func RunServer() {
	router := mux.NewRouter()

	for _, route := range Routes {
		fmt.Printf("⚙️   The %s %s route is added.\n", route.method, route.pattern)
		router.HandleFunc(route.pattern, WrapHandler(route.handler))
	}

	fmt.Printf("✅  The server is running at http://localhost:%d.\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
