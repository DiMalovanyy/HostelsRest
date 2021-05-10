package restapi

import (
	"database/sql"
	"net/http"

	"github.com/UniverOOP/internal/app/store/postgresStore"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := postgresStore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))

	serv, err := NewServer(config.LogLevel, store, sessionStore)
	if err != nil {
		return err
	}
	handler := cors.Default().Handler(serv)
	return http.ListenAndServe(config.BindAddress, handler)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
