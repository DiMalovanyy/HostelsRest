package restapi

import (
	"net/http"

	"github.com/UniverOOP/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
}

func NewServer(logLevel string, store store.Store) (*server, error) {
	logger, err := configureLogger(logLevel)
	if err != nil {
		return nil, err
	}

	s := &server{
		router: mux.NewRouter(),
		logger: logger,
	}

	s.configureRouter()
	logger.Info("Server started")
	return s, err
}

func configureLogger(logLevel string) (*logrus.Logger, error) {
	logLvl, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.SetLevel(logLvl)
	return logger, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/register", s.handlerRegisterRequest()).Methods("POST")
	s.router.HandleFunc("/login", s.handlerLoginRequest()).Methods("POST")

	s.router.HandleFunc("/faculties", s.handlerFacultiesRequest()).Methods("GET")
	s.router.HandleFunc("/hostels", s.handlerHostelsRequest()).Methods("GET")

	//When user authed
	s.router.HandleFunc("/upgrade_user", s.handleUpgradeUserRequest()).Methods("POST")
}

func (s *server) handlerLoginRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) handlerRegisterRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) handlerFacultiesRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) handlerHostelsRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}

func (s *server) handleUpgradeUserRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

	}
}
