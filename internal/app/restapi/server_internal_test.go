package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/UniverOOP/internal/app/store/testStore"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/register", nil)
	s, err := NewServer("error", testStore.New())

	assert.NoError(t, err)
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
