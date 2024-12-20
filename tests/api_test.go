package tests

import (
	"bytes"
	"coditas-task/handlers"
	"coditas-task/validators"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_ValidData(t *testing.T) {
	// Create a new Gin instance for the test
	r := gin.Default()

	// Initialize the validator and register custom validation
	validate := validator.New()
	validators.RegisterValidators(validate)

	// Initialize the UserHandler with the validator
	userHandler := handlers.NewUserHandler(validate)

	// POST route for the user creation
	r.POST("/user", userHandler.CreateUser)

	// Valid payload
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE1234F",
		"mobile": "1234567890",
		"email":  "john.doe@example.com",
	}

	// Marshal the user payload to JSON
	jsonValue, _ := json.Marshal(user)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response using httptest.NewRecorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response status and message
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User created successfully")
}

func TestCreateUser_InvalidPAN(t *testing.T) {
	r := gin.Default()

	validate := validator.New()
	validators.RegisterValidators(validate)

	userHandler := handlers.NewUserHandler(validate)

	r.POST("/user", userHandler.CreateUser)

	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE12345F", // Invalid PAN (should be ABCDE1234F)
		"mobile": "1234567890",
		"email":  "john.doe@example.com",
	}

	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "PAN")
}

func TestCreateUser_InvalidMobile(t *testing.T) {
	// Create a new Gin instance for the test
	r := gin.Default()

	// Initialize the validator and register custom validation
	validate := validator.New()
	validators.RegisterValidators(validate)

	// Initialize the UserHandler with the validator
	userHandler := handlers.NewUserHandler(validate)

	// POST route for the user creation
	r.POST("/user", userHandler.CreateUser)

	// Invalid mobile number (less than 10 digits)
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE1234F",
		"mobile": "12345", // Invalid mobile (should be exactly 10 digits)
		"email":  "john.doe@example.com",
	}

	// Marshal the user payload to JSON
	jsonValue, _ := json.Marshal(user)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response using httptest.NewRecorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response status and check the error message
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Mobile")
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	// Create a new Gin instance for the test
	r := gin.Default()

	// Initialize the validator and register custom validation
	validate := validator.New()
	validators.RegisterValidators(validate)

	// Initialize the UserHandler with the validator
	userHandler := handlers.NewUserHandler(validate)

	// POST route for the user creation
	r.POST("/user", userHandler.CreateUser)

	// Invalid email (wrong format)
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE1234F",
		"mobile": "1234567890",
		"email":  "john.doe@com", // Invalid email (should be a valid email address)
	}

	// Marshal the user payload to JSON
	jsonValue, _ := json.Marshal(user)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Record the response using httptest.NewRecorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response status and check the error message
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Email")
}
