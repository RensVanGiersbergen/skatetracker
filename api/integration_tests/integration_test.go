package integration_tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/RensVanGiersbergen/skatetracker/repository"
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
	assert.Equal(t, `"chill"`, w.Body.String())
}

func TestLoginRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"email":"test@gmail.com", "password":"admin"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/login", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user logged in")
}

func TestLoginRoute_UnusedEmail(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"email":"notauser@gmail.com", "password":"admin"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/login", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, `{"error":"user not found"}`, w.Body.String())
}

func TestLoginRoute_InvalidCredentials(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"email":"test@gmail.com", "password":"notadmin"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/login", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, `{"error":"invalid credentials"}`, w.Body.String())
}

func TestRegisterRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"username":"test2", "email":"test2@gmail.com", "password":"test"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/register", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"user registered"}`, w.Body.String())
}

func TestRegisterRoute_InvalidEmail(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"username":"test3", "email":"notanemail", "password":"test"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/register", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"invalid email"}`, w.Body.String())
}

func TestRegisterRoute_DuplicateUsername(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"username":"test1", "email":"test4@gmail.com", "password":"test"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/register", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"username already in use"}`, w.Body.String())
}

func TestRegisterRoute_DuplicateEmail(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Create a new request with a JSON body
	jsonBody := `{"username":"test5", "email":"test@gmail.com", "password":"test"}`
	reqBody := strings.NewReader(jsonBody)

	// Act
	req, _ := http.NewRequest("POST", "/account/register", reqBody)
	req.Header.Set("Content-type", "application/json")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"email already in use"}`, w.Body.String())
}

func TestVerifyRoute_Success(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops
	os.Setenv("JWT_SECRET", "8xp4faT3fluB7wFhn2Bevbblk4OGOfuufIxKi9KRDZvnzH2Jwj9GGuukX6HTVgr")

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/account/verify", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNzY0MDI3MTgzLCJzdWIiOiI3ZjczMzIwMC1kN2QzLTRmNTYtYjI0Ny1mZjlmMWQ5NTdlZjAiLCJ1c2VybmFtZSI6InRlc3QxIn0.i2jCrIB5fATnMLa4lwTSF8DA8X88KE6F_ZHs3VneYHA")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user is verified")
}

func TestVerifyRoute_WrongSecretKey(t *testing.T) {
	// Arrange
	// Set up database connection
	connstr := "host=localhost port=5432 user=postgres password=pizza123 dbname=skatetrackerdev sslmode=disable"
	repository.InitPostgresDB(connstr)
	_ = repository.LoadQueries()
	defer repository.ClosePostgresDB() // Ensure the DB connection is closed when the app stops
	os.Setenv("JWT_SECRET", "someotherkey")

	// Set up router
	router := router.SetupRouter()
	w := httptest.NewRecorder()

	// Act
	req, _ := http.NewRequest("GET", "/account/verify", nil)
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZ21haWwuY29tIiwiZXhwIjoxNzY0MDI3MTgzLCJzdWIiOiI3ZjczMzIwMC1kN2QzLTRmNTYtYjI0Ny1mZjlmMWQ5NTdlZjAiLCJ1c2VybmFtZSI6InRlc3QxIn0.i2jCrIB5fATnMLa4lwTSF8DA8X88KE6F_ZHs3VneYHA")
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, `{"error":"unauthorized"}`, w.Body.String())
}
