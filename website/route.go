package website

import (
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	method  string
	pattern string
	handler ParametrizedHttpHandlerFunc
}

type HttpParameters = map[string]string
type ParametrizedHttpHandlerFunc = func(response http.ResponseWriter, request *http.Request, parameters HttpParameters)

func createRoute(method string, pattern string, handler ParametrizedHttpHandlerFunc) route {
	return route{method, pattern, handler}
}

func Get(pattern string, handler ParametrizedHttpHandlerFunc) route {
	return createRoute("GET", pattern, handler)
}

func WrapHandler(handler ParametrizedHttpHandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var parameters HttpParameters = mux.Vars(request)
		handler(response, request, parameters)
	}
}
