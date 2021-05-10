package restapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/UniverOOP/internal/app/model"
	"github.com/UniverOOP/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName = "univerApi"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func NewServer(logLevel string, store store.Store, sessionStore sessions.Store) (*server, error) {
	logger, err := configureLogger(logLevel)
	if err != nil {
		return nil, err
	}

	s := &server{
		router:       mux.NewRouter(),
		logger:       logger,
		store:        store,
		sessionStore: sessionStore,
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

	s.router.HandleFunc("/faculty_hostels", s.handlerFacultyHostles()).Methods("GET")

	//When user authed
	s.router.HandleFunc("/upgrade_user", s.handleUpgradeUserRequest()).Methods("POST")
}

func (s *server) handlerLoginRequest() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(rw, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
			return
		}
		session.Values["user_id"] = u.Id
		if err := s.sessionStore.Save(r, rw, session); err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(rw, r, http.StatusOK, nil)
	}
}

func (s *server) handlerRegisterRequest() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {

		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().CreateUser(u); err != nil {
			s.error(rw, r, http.StatusUnprocessableEntity, err)
			return
		}
		u.Sanitize()
		s.respond(rw, r, http.StatusCreated, u)
	}
}

func (s *server) handlerFacultyHostles() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {

		type ResponseInnerStruct struct {
			hostel_name string `json:"hostel_name"`
		}
		type ResponseStruct struct {
			faculty_name string                `json:"faculty_name"`
			hostels      []ResponseInnerStruct `json:"hostels"`
		}

		responseStruct := make([]ResponseStruct, 0)

		faculties, err := s.store.Faculty().GetAllFaculties()
		if err != nil {
			s.error(rw, r, http.StatusUnprocessableEntity, err)
		}

		for _, fac := range faculties {
			hostels, err := s.store.Hostel().GetHostelsByFucultyId(fac.Id)
			if err != nil {
				continue
			}
			hostelsStr := make([]ResponseInnerStruct, 0)
			for _, hs := range hostels {
				hostelsStr = append(hostelsStr, ResponseInnerStruct{hostel_name: hs.Description})
			}
			responseStruct = append(responseStruct, ResponseStruct{faculty_name: fac.Name, hostels: hostelsStr})
		}
		log.Print(responseStruct)
		s.respond(rw, r, http.StatusOK, responseStruct)
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

func (s *server) error(rw http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(rw, r, code,
		map[string]string{"error": err.Error()},
	)
}

func (s *server) respond(rw http.ResponseWriter, r *http.Request, code int, data interface{}) {
	rw.WriteHeader(code)

	if data != nil {
		json.NewEncoder(rw).Encode(data)
	}
}
