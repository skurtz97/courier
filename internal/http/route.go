package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Route struct {
	Method      string       `json:"method"`
	Path        string       `json:"path"`
	Description string       `json:"description"`
	UrlParam    UrlParam     `json:"url_param,omitempty"`
	QueryParams []QueryParam `json:"query_params,omitempty"`
	Body        interface{}  `json:"body,omitempty"`
}

type UrlParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type QueryParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// ReadRoutesFile reads the routes definition file ("routes.json" at project root) and returns a slice of routes.
func ReadRoutesFile(filePath string) ([]Route, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var routes []Route
	err = decoder.Decode(&routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

// List routes returns a json response with all the routes.
func (s *Server) ListRoutes(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(s.routes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s.log.Error(err.Error())
		w.Write([]byte(`{"message": "internal error listing routes"}`))
	}
	w.Write(js)
}
