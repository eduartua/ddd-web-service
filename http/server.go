package http

import (
	"net/http"

	app "github.com/eduartua/ddd-web-service"
	"github.com/gorilla/mux"
)

func NewServer(
	us app.UserStore,
	// Bytes to be used my middleware when implemented
	bytes []byte,
) http.Handler {
	html := HTMLServer(us)
	json := JSONServer(us)
	m := http.NewServeMux()
	m.Handle("/", html)
	m.Handle("/api/", http.StripPrefix("/api", json))
	return m
}

func HTMLServer(us app.UserStore) http.Handler {
	router := mux.NewRouter()
	svr := server {
		router: router,
		static: htmlStaticHandler(),
		users: htmlUserHandler(us),
	}
	svr.routes(true)
	return &svr
}

func JSONServer(us app.UserStore) http.Handler {
	svr := server{
		users: jsonUserHandler(us),
	}
	svr.routes(false)
	return &svr
}

type server struct {
	router *mux.Router
	static *StaticHandler
	users  *UserHandler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) routes(webMode bool) {
	if webMode {
		s.router.HandleFunc("/", s.static.Home).Methods("GET")
	}
}