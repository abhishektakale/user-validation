package tests

import (
	"bytes"
	"coditas-task/handlers"
	"coditas-task/validators"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = gin.Default()

	validate := validator.New()
	validators.RegisterValidators(validate)

	userHandler := handlers.NewUserHandler(validate)

	router.POST("/user", userHandler.CreateUser)

	code := m.Run()

	os.Exit(code)
}

// Helper function to send a POST request
func sendPostRequest(user map[string]interface{}) *httptest.ResponseRecorder {
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	return w
}

func TestCreateUser_ValidData(t *testing.T) {
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE1234F",
		"mobile": "1234567890",
		"email":  "john.doe@example.com",
	}

	w := sendPostRequest(user)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User created successfully")
}

func TestCreateUser_InvalidPAN(t *testing.T) {
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE12345F", // Invalid PAN (should be ABCDE1234F)
		"mobile": "1234567890",
		"email":  "john.doe@example.com",
	}

	w := sendPostRequest(user)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "PAN")
}

func TestCreateUser_InvalidMobile(t *testing.T) {
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE1234F",
		"mobile": "12345", // Invalid mobile (should be exactly 10 digits)
		"email":  "john.doe@example.com",
	}

	w := sendPostRequest(user)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Mobile")
}

func TestCreateUser_InvalidEmail(t *testing.T) {
	user := map[string]interface{}{
		"name":   "John Doe",
		"pan":    "ABCDE1234F",
		"mobile": "1234567890",
		"email":  "john.doe@com", // Invalid email (should be a valid email address)
	}

	w := sendPostRequest(user)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Email")
}
