package handler

import (
	"net/http"

	"Golang-Redis-Gin/internal/controllers"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    ctrl controllers.UserController
}

func NewUserHandler(c controllers.UserController) *UserHandler {
    return &UserHandler{ctrl: c}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
    id := c.Param("id")
    user, err := h.ctrl.GetUserByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
        return
    }

    c.JSON(http.StatusOK, user)
}
