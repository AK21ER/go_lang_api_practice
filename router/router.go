package router

import (
	"net/http"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type Route struct {
	Method  string
	Path    string
	Handler HandlerFunc
}

var routes []Route


func GET(path string, handler HandlerFunc) {
	routes = append(routes, Route{
		Method:  "GET",
		Path:    path,
		Handler: handler,
	})
}

func POST(path string, handler HandlerFunc) {
	routes = append(routes, Route{
		Method:  "POST",
		Path:    path,
		Handler: handler,
	})
}

func DELETE(path string, handler HandlerFunc) {
	routes = append(routes, Route{
		Method:  "DELETE",
		Path:    path,
		Handler: handler,
	})
}

func PUT(path string, handler HandlerFunc) {
	routes = append(routes, Route{
		Method:  "PUT",
		Path:    path,
		Handler: handler,
	})
}

func Serve(w http.ResponseWriter, r *http.Request) {

	for _, route := range routes {

		if r.Method != route.Method {
			continue //continue means: "skip the rest of this loop iteration and go to the next route"
		}

		params, ok := match(route.Path, r.URL.Path)

		if ok {
			route.Handler(w, r, params)
			return
		}
	}

	http.Error(w, "Route not found", http.StatusNotFound)
}

func match(routePath, requestPath string) (map[string]string, bool) {

	routeParts := strings.Split(routePath, "/")
	reqParts := strings.Split(requestPath, "/")

	if len(routeParts) != len(reqParts) {
		return nil, false
	}

	params := make(map[string]string)

	for i := range routeParts {

		if strings.HasPrefix(routeParts[i], ":") {
			paramName := routeParts[i][1:]
			params[paramName] = reqParts[i]
		} else if routeParts[i] != reqParts[i] {
			return nil, false
		}
	}

	return params, true
}