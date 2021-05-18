package restapi

import (
	"context"
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
	sessionName        = "univerApi"
	ctxKeyUser  ctxKey = iota
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
	errHostelsNotFound          = errors.New("hostels for faulty not found")
	errNoFreeRooms              = errors.New("no free rooms in faculty")
	errIncorrectSex             = errors.New("incorrect request sex")
	// errUserNotSetteled          = errors.New("user not settled")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

//Midleware
func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().FindById(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
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

	s.router.HandleFunc("/faculty_hostels", s.handlerFacultyHostles()).Methods("GET")
	s.router.HandleFunc("/faculties", s.handlerGetAllFaculties()).Methods("GET")
	s.router.HandleFunc("/user_status", s.handleGetUserStatus()).Methods("GET")

	//When user authed
	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/logout", s.handlerLogoutRequest()).Methods("GET")
	private.HandleFunc("/upgrade_user", s.handleUpgradeUserRequest()).Methods("POST")
	private.HandleFunc("/hostel_room_members", s.handleHostelRoomMembers()).Methods("GET")
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

		rw.Header().Add("access-control-expose-headers", "Set-Cookie")
		s.respond(rw, r, http.StatusOK, nil)
	}
}

func (s *server) handlerLogoutRequest() http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {
		sessions, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
			return
		}

		sessions.Options.MaxAge = -1
		if err := s.sessionStore.Save(r, rw, sessions); err != nil {
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
		type Hostel struct {
			Hostel_name string `json:"hostel_name"`
		}
		type Faculty struct {
			Faculty_name string   `json:"faculty_name"`
			Housings     []Hostel `json:"housings"`
		}

		type Response []Faculty

		response := make(Response, 0)

		faculties, err := s.store.Faculty().GetAllFaculties()
		if err != nil {
			s.error(rw, r, http.StatusUnprocessableEntity, err)
		}

		for _, fac := range faculties {
			hostels, err := s.store.Hostel().GetHostelsByFucultyId(fac.Id)
			if err != store.ErrEmptyData && err != nil {
				continue
			}

			hostelsStr := make([]Hostel, 0)
			for _, hs := range hostels {
				hostelsStr = append(hostelsStr, Hostel{Hostel_name: hs.Description})
			}
			response = append(response, Faculty{Faculty_name: fac.Name, Housings: hostelsStr})
		}

		s.respond(rw, r, http.StatusOK, response)
	}
}

func (s *server) handlerGetAllFaculties() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fucs := make([]string, 0)
		faculties, err := s.store.Faculty().GetAllFaculties()
		if err != nil {
			s.error(rw, r, http.StatusUnprocessableEntity, err)
		}

		for _, f := range faculties {
			fucs = append(fucs, f.Name)
		}

		s.respond(rw, r, http.StatusOK, fucs)
	}
}

func (s *server) handleGetUserStatus() http.HandlerFunc {

	return func(rw http.ResponseWriter, r *http.Request) {

		var responseBool bool
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(rw, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().FindById(id.(int))
		if err != nil {
			s.error(rw, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		if u.RoomId == 0 || u.FacultyId == 0 {
			responseBool = false
		} else {
			responseBool = true
		}
		s.respond(rw, r, http.StatusOK, responseBool)
	}
}

func (s *server) handleHostelRoomMembers() http.HandlerFunc {

	type Room struct {
		RoomNumber string   `json:"room_number"`
		Names      []string `json:"names"`
	}

	type Response struct {
		HostelName string `json:"hostel_name"`
		Rooms      []Room `json:"rooms"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {

		//Get current session user
		currentUser, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(rw, r, http.StatusInternalServerError, nil)
		}

		log.Print("Room id: ", currentUser.RoomId)
		room, err := s.store.Room().GetRoomByRoomId(currentUser.RoomId)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
		}

		rooms, err := s.store.Room().GetAllRoomsByHostleId(room.HostelId)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
		}

		hostel, err := s.store.Hostel().GetHostelById(room.HostelId)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
		}

		// log.Print("Hostel: ", hostel.Description)

		var response Response
		response.HostelName = hostel.Description

		for _, room = range rooms {
			var currentRoom Room
			currentRoom.Names = make([]string, 0)

			currentRoom.RoomNumber = room.Number
			users, err := s.store.User().GetAllUsersByRoomId(room.Id)
			if err != nil && err != store.ErrEmptyData {
				s.error(rw, r, http.StatusInternalServerError, err)
			}
			for _, user := range users {
				currentRoom.Names = append(currentRoom.Names, user.Name)
			}
			response.Rooms = append(response.Rooms, currentRoom)
		}

		s.respond(rw, r, http.StatusOK, response)
	}
}

//test it
func (s *server) handleUpgradeUserRequest() http.HandlerFunc {

	type request struct {
		DegreeLevel int    `json:"degreeLevel"`
		Sex         string `json:"sex"`
		FacultyName string `json:"facultyName"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(rw, r, http.StatusBadRequest, err)
			return
		}

		var sex model.Sex
		if req.Sex == "male" {
			sex = model.MEN
		} else if req.Sex == "female" {
			sex = model.WOMEN
		} else {
			s.error(rw, r, http.StatusBadRequest, errIncorrectSex)
		}

		// log.Printf("sex: %d", sex)

		faculty, err := s.store.Faculty().GetFacultyByName(req.FacultyName)
		if err != nil {
			s.error(rw, r, http.StatusInternalServerError, err)
		}

		//Get current session user
		currentUser, ok := r.Context().Value(ctxKeyUser).(*model.User)
		if !ok {
			s.error(rw, r, http.StatusInternalServerError, nil)
		}

		hostels, err := s.store.Hostel().GetHostelsByFucultyId(faculty.Id)
		if err != nil {
			if err == store.ErrEmptyData {
				s.error(rw, r, http.StatusNoContent, errHostelsNotFound)
			} else {
				s.error(rw, r, http.StatusInternalServerError, err)
			}
		}

		isFound := false
		for _, hostel := range hostels {

			roomId, err := s.store.Room().GetFreeRoomByHostelId(hostel.Id)
			if err != nil {
				if err == store.ErrNoData || err == store.ErrEmptyData {
					continue
				} else {
					s.error(rw, r, http.StatusInternalServerError, err)
				}
			}
			isFound = true
			if err := s.store.User().Upgrade(
				currentUser.Id,
				sex, roomId,
				faculty.Id, req.DegreeLevel); err != nil {
				s.error(rw, r, http.StatusInternalServerError, err)
			}
		}

		if !isFound {
			s.error(rw, r, http.StatusNoContent, errNoFreeRooms)
		}

		s.respond(rw, r, http.StatusOK, nil)
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
