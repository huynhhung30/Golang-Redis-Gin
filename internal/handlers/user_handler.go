package handler

import (
	"Golang-Redis-Gin/internal/constants"
	"Golang-Redis-Gin/internal/controllers"
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/utils"
	"Golang-Redis-Gin/internal/utils/functions"

	"net/http"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
    ctrl *controllers.UserController
}

func NewUserHandler(ctrl *controllers.UserController) *UserHandler {
    return &UserHandler{ctrl: ctrl}
}

func (h *UserHandler) LogInUser(c *gin.Context) {
	req := &models.UserModel{}
	functions.ShowLog("&models.UserModel{}", req)

   if err := c.ShouldBindJSON(&req); err != nil {
    utils.Error(c, http.StatusBadRequest, "Yêu cầu không hợp lệ")
    return
}
functions.ShowLog("req", req)
    user, err := h.ctrl.GetUserByEmail(c.Request.Context(), req.Email)
    	if err != nil || !utils.CheckPassword(req.Password, user.Password) {
		utils.Error(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}
 
	accessToken, err := utils.GenerateAccessToken(user.Id, user.Email, user.Role)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate access token")
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Id)
	if err != nil {
		utils.Error(c, http.StatusInternalServerError, "Failed to generate refresh token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
    
}
func (h *UserHandler) CreateUser(c *gin.Context) {
    var req models.UserModel
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.Error(c, http.StatusBadRequest, "Invalid request body")
        return
    }

    user := &models.UserModel{
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Email:     req.Email,
        Password:  req.Password,
    }

    user, err := h.ctrl.Create(c.Request.Context(), user)
    if err != nil {
        if appErr, ok := err.(*utils.AppError); ok {
            utils.Error(c, appErr.Code, appErr.Message)
            return
        }
        utils.Error(c, http.StatusInternalServerError, "Internal server error")
        return
    }

    resp := gin.H{
        "id":        user.Id,
        "firstName": user.FirstName,
        "lastName":  user.LastName,
        "email":     user.Email,
    }

    utils.Created(c, constants.SUCCESS, resp)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	user, err := h.ctrl.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}