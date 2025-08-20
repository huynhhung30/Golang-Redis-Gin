package handler

import (
	"Golang-Redis-Gin/internal/controllers"
	"Golang-Redis-Gin/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
    ctrl *controllers.UserController
}

func NewUserHandler(ctrl *controllers.UserController) *UserHandler {
    return &UserHandler{ctrl: ctrl}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var user models.UserModel
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
        return
    }
    if err := h.ctrl.Create(c.Request.Context(), &user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
    id := c.Param("id")
    user, err := h.ctrl.GetByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}