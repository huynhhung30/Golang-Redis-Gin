package controllers

import (
	"net/http"
	"trinity-app/config"
	"trinity-app/models"
	"trinity-app/utils"
	"trinity-app/utils/constants"
	"trinity-app/utils/functions"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)




func Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

// MemberRegister 
// @Summary Create a user
// @Description Create a new user with the given input data
// @Tags user
// @Accept  json
// @Produce  json
// @Param   input    body  models.UserModel true "Require Email and Password"
// @Success 200
// @Failure 406
// @Router /member-register [post]
func MigrateModel(c *gin.Context) {
	config.DB.Debug().AutoMigrate(
		models.UserModel{},
		models.CouponModel{},
		)
	RES_SUCCESS_SIMPLE(c, "Create model sucessfully")
}
// MemberRegister 
// @Summary Create a user
// @Description Create a new user with the given input data
// @Tags user
// @Accept  json
// @Produce  json
// @Param   input    body  models.UserModel true "Require Email and Password"
// @Success 200
// @Failure 406
// @Router /member-register [post]
func MemberRegister(c *gin.Context) {
	// Create models
	requestBody := models.UserModel{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.UserType = constants.USER_TYPE_MEMBER
	requestBody.LoginMethod = constants.LOGIN_METHOD_SYSTEM
	requestBody.IsSilver = false
	// Check input empty
	if requestBody.Password == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "password is required")
		return
	}
	if requestBody.Email == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "Email is required")
		return
	}
	// Check exist email
	us := models.FindUserProfileByEmail(requestBody.Email, constants.USER_TYPE_MEMBER)
	if us.Id != 0 {
		RES_ERROR_MSG(c, http.StatusConflict, "This email already exists", nil)
		return
	}
	// Hash password
	requestBody.Password = utils.HashPassword(requestBody.Password)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	userInfo, err := models.CreateUser(&requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	RES_SUCCESS_MSG(c, userInfo, "Register member sucessfully!")
}


// @BasePath /api/v1/trinity
// LoginHandler 
// @Summary Silver Register
// @Schemes
// @Description Require input Id of coupon 
// @Tags user
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param   id_coupon     body    models.QueryCoupon     true        "POST"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /silver-register [post]
func SilverRegister(c *gin.Context) {
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_USER_NOT_FOUND, nil)
		return
	}
	userInfo := models.FindUserProfileById(tokenInfo.UserId)
	if userInfo.IsSilver != false{
		RES_ERROR_MSG(c, http.StatusConflict, "You are a silver member", nil)
		return
	}
	requestBody := models.QueryCoupon{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	if requestBody.Id == 0 {
		models.SilverRegisterMember(tokenInfo.UserId)
		RES_SUCCESS_MSG(c, "", "Register member sucessfully!")
	}else{
		couponInfo := models.FindCouponId(requestBody.Id)
		now := functions.CurrentTime()
		if functions.InTimeSpan(couponInfo.CreatedAt, couponInfo.DueDate, now) == false || couponInfo.IsValid != true {
			RES_ERROR_MSG(c, http.StatusConflict, "Coupon Expired ", nil)
			return
		}
		models.SilverRegisterMember(tokenInfo.UserId)
		models.UpdateCountRegistrations(couponInfo)
		RES_SUCCESS_MSG(c, couponInfo, "Register member discount 30% sucessfully! ")
	}
}





// Member Register Social
func MemberRegisterSocial(c *gin.Context) {
	// Create models
	requestBody := models.UserModel{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.UserType = constants.USER_TYPE_MEMBER
	// Check input empty
	if requestBody.LoginMethod == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "login_method is required")
		return
	}
	if requestBody.LoginMethod != constants.LOGIN_METHOD_GOOLE && requestBody.LoginMethod != constants.LOGIN_METHOD_FACEBOOK && requestBody.LoginMethod != constants.LOGIN_METHOD_APPLE {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "login_method must be: google, facebook, apple")
		return
	}
	if requestBody.SocialId == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "social_id is required")
		return
	}
	if requestBody.LastName == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "last_name is required")
		return
	}
	// Check user exist
	userInfoExist := models.FindUserProfileBySocialId(requestBody.LoginMethod, requestBody.SocialId)
	if userInfoExist.Id != 0 {
		RES_SUCCESS_MSG(c, userInfoExist, "Register member social sucessfully!")
		return
	}
	userInfo, err := models.CreateUser(&requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	RES_SUCCESS_MSG(c, userInfo, "Register member social sucessfully!")
}

// Member Login

// @BasePath /api/v1/trinity
// LoginHandler 
// @Summary Authenticate a user
// @Schemes
// @Description Authenticates a user using email and password, returns a JWT token if successful
// @Tags user
// @Accept  json
// @Produce  json
// @Param   user     body    models.UsersLoginRequestModel     true        "User login object"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /member-login [post]
func MemberLogin(c *gin.Context) {
	requestBody := models.UsersLoginRequestModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.LoginMethod = constants.LOGIN_METHOD_SYSTEM
	// Check input empty
	if requestBody.Email == "" || requestBody.Password == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, "email and password is required", err)
		return
	}
	// Find user by email
	userInfo := models.FindUserProfileByEmail(requestBody.Email, constants.USER_TYPE_MEMBER)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_EMAIL_INCORRECT, nil)
		return
	}
	// Check password
	isMathPassword := utils.CheckPasswordHash(requestBody.Password, userInfo.Password)
	if isMathPassword == false {
		RES_ERROR_MSG(c, http.StatusMethodNotAllowed, constants.MSG_PASSWORD_INCORRECT, nil)
		return
	}
	// Generate token
	userInfo.Token = utils.GenerateTokenString(userInfo.Id, constants.USER_TYPE_MEMBER)
	userInfo.Password = ""
	RES_SUCCESS(c, userInfo)
}

// Member Login Social
func MemberLoginSocial(c *gin.Context) {
	requestBody := models.UsersLoginRequestModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	// Check input empty
	if requestBody.LoginMethod == "" || requestBody.SocialId == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, "email and password is required", err)
		return
	}
	if requestBody.LoginMethod != constants.LOGIN_METHOD_GOOLE && requestBody.LoginMethod != constants.LOGIN_METHOD_FACEBOOK && requestBody.LoginMethod != constants.LOGIN_METHOD_APPLE {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, "login_method must be: google, facebook, apple")
		return
	}
	// Find user by socical id
	userInfo := models.FindUserProfileBySocialId(requestBody.LoginMethod, requestBody.SocialId)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusNotFound, constants.MSG_USER_NOT_FOUND, nil)
		return
	}
	// Generate token
	userInfo.Token = utils.GenerateTokenString(userInfo.Id, constants.USER_TYPE_MEMBER)
	RES_SUCCESS(c, userInfo)
}

// Admin Login
func AdminLogin(c *gin.Context) {
	requestBody := models.UsersLoginRequestModel{}
	err := BindJSON(c, &requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}
	requestBody.LoginMethod = constants.LOGIN_METHOD_SYSTEM
	// Check input empty
	if requestBody.Email == "" || requestBody.Password == "" {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, "email and password is required", err)
		return
	}
	// Find user by email
	userInfo := models.FindUserProfileByEmail(requestBody.Email, constants.USER_TYPE_ADMIN)
	if userInfo.Id == 0 {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_EMAIL_INCORRECT, nil)
		return
	}
	// Check password
	isMathPassword := utils.CheckPasswordHash(requestBody.Password, userInfo.Password)
	if isMathPassword == false {
		RES_ERROR_MSG(c, http.StatusMethodNotAllowed, constants.MSG_PASSWORD_INCORRECT, nil)
		return
	}
	// Generate token
	userInfo.Token = utils.GenerateTokenString(userInfo.Id, constants.USER_TYPE_ADMIN)
	userInfo.Password = ""
	RES_SUCCESS(c, userInfo)
}
