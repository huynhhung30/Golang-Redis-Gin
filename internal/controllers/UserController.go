package controllers

import (
	"Golang-Redis-Gin/internal/models"
	"Golang-Redis-Gin/internal/service"
	"Golang-Redis-Gin/internal/utils"
	"Golang-Redis-Gin/internal/utils/constants"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)



type UserController interface {
    GetUserByID(ctx context.Context, id string) (*models.UserModel, error)
}

type userController struct {
    svc service.UserService
}

func NewUserController(s service.UserService) UserController {
    return &userController{svc: s}
}

func (c *userController) GetUserByID(ctx context.Context, id string) (*models.UserModel, error) {
    return c.svc.GetProfile(ctx, id)
}


// Get profile user
// func GetUserProfile(c *gin.Context) {
// 	tokenInfo := utils.GetTokenInfo(c)
// 	if tokenInfo.UserId == 0 {
// 		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
// 		return
// 	}
// 	userInfo := models.FindUserProfileById(tokenInfo.UserId)
// 	RES_SUCCESS(c, userInfo)
// }

// GetUserProfileById

// func GetUserProfileById(c *gin.Context) {
// 	// tokenInfo := utils.GetTokenInfo(c)
// 	// if tokenInfo.UserId == 0 {
// 	// 	RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
// 	// 	return
// 	// }
// 	idParam :=  c.Param("id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_INVALID_INPUT, nil)
// 		return
// 	}
// 	userInfo := models.FindUserProfileById(id)
// 	if userInfo.Id == 0 {
// 		RES_SUCCESS_SIMPLE(c, nil)
// 	} else {
// 		RES_SUCCESS(c, userInfo)
// 	}
// }

// Update Profile
func UpdateProfile(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	userUpdateBody := models.UserModel{}
	if err := c.ShouldBindBodyWith(&userUpdateBody, binding.JSON); err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, constants.MSG_INVALID_INPUT, err)
		return
	}
	userInfo := models.FindUserInfoById(tokenInfo.UserId)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_USER_NOT_FOUND, nil)
		return
	}
	userUpdateBody.Id = tokenInfo.UserId
	userInfo, err := models.UpdateUser(userUpdateBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, userInfo, "Update profile successfully")
}

// Update Fcm Token
func UpdateFcmToken(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	fcmTokenBody := models.FcmTokenModel{}
	if err := c.ShouldBindBodyWith(&fcmTokenBody, binding.JSON); err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, constants.MSG_INVALID_INPUT, err)
		return
	}
	fcmTokenInfo := models.FindFcmToken(fcmTokenBody.FcmToken)
	if fcmTokenInfo.Id != 0 {
		RES_SUCCESS_MSG(c, fcmTokenInfo, "Update profile successfully")
		return
	}
	fcmTokenBody.Id = tokenInfo.UserId
	fcmTokenInfoNew, err := models.UpdateFcmToken(fcmTokenBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotFound, err.Error(), nil)
		return
	}
	RES_SUCCESS_MSG(c, fcmTokenInfoNew, "Update fcm token successfully")
}
