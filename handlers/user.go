package handlers

import (
	"coditas-task/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserHandler defines the handler with a validator instance
type UserHandler struct {
	Validator *validator.Validate
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(v *validator.Validate) *UserHandler {
	if v == nil {
		panic("validator instance cannot be nil")
	}

	return &UserHandler{Validator: v}
}

// CreateUser handles the POST request for creating a new user
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	if err := h.Validator.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
