package postgresStore

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABABSE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5434 dbname=univerTest user=postgres password=12345 sslmode=disable"
	}

	os.Exit(m.Run())
}
