package restapi

import (
	"database/sql"
	"net/http"

	"github.com/UniverOOP/internal/app/store/postgresStore"
	_ "github.com/lib/pq"
)

func Start(config *Config) error {

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	serv, err := NewServer(config.LogLevel, postgresStore.New(db))
	if err != nil {
		return err
	}
	return http.ListenAndServe(config.BindAddress, serv)
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
