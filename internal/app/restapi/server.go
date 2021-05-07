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
	s.router.HandleFunc("/hello", s.handlerHelloRequest())
}

func (s *server) handlerHelloRequest() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello"))
	}
}
