package controller

import (
	"net/http"
	"test-naga-exchange/model"
	"test-naga-exchange/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	if err := h.service.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Register success"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	u, err := h.service.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "invalid credentials"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Login success", "data": gin.H{"token": u.Token}})
}
