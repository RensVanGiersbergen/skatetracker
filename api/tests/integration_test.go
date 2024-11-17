package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RensVanGiersbergen/skatetracker/router"
	"github.com/stretchr/testify/assert"
)

func TestTestRoute_Success(t *testing.T) {
	// Arrange
	r := router.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `"Chill"`, w.Body.String())
}
