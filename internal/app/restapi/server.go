package restapi

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func NewServer() *server {

	s := &server{
		router: mux.NewRouter(),
	}

	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/hello", s.handlerHelloRequest())
}

func (s *server) handlerHelloRequest() http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	}
}
