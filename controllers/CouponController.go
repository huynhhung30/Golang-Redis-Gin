package controllers

import (
	"net/http"
	"Golang-Redis-Gin/models"
	"Golang-Redis-Gin/utils"
	"Golang-Redis-Gin/utils/constants"
	"Golang-Redis-Gin/utils/functions"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary Create a coupon
// @Description Create a new coupon
// @Tags Coupon
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param   input    body  models.CouponModel true "Require Email and Password"
// @Success 200
// @Failure 406
// @Router /create-coupon [post]
func CreateCoupon(c *gin.Context) {
	// Create models
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_USER_NOT_FOUND, nil)
		return
	}
	requestBody := models.CouponModel{}
	err := c.ShouldBindBodyWith(&requestBody, binding.JSON)
	functions.ShowLog("requestBody", requestBody)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err)
		return
	}

	requestBody.UserId = tokenInfo.UserId
	rs, err := models.CreateCoupon(&requestBody)
	functions.ShowLog("rs", rs)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusNotAcceptable, constants.MSG_INVALID_INPUT, err.Error())
		return
	}
	
	RES_SIMPLE(c,"Create Coupon sucessfully!")
}

// Get User Profile List


// GetCouponList 
// @Get Coupon List
// @Schemes
// @Description Get Coupon List
// @Tags Coupon
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param   coupon     query    models.PageLimitQueryModel     flase      "get coupon list"
// @Success 200 {array} array "list coupon"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /coupon-list [get]
func GetCouponList(c *gin.Context) {
	queryParams := models.PageLimitQueryModel{}
	err := c.BindQuery(&queryParams)
	if err != nil {
		RES_ERROR_MSG(c, http.StatusForbidden, constants.MSG_INVALID_INPUT, err)
		return
	}
	if queryParams.Limit == 0 {
		queryParams.Limit = constants.LIMIT_DEFAULT
	}
	tokenInfo := utils.GetTokenInfo(c)
	if tokenInfo.UserId == 0 {
		RES_ERROR_MSG(c, http.StatusUnauthorized, constants.MSG_TOKEN_NOT_FOUND, nil)
		return
	}
	userList, totalCount := models.FindCouponList(queryParams)
	meta := GeneralPaginationModel{}
	meta.CurrentPage = queryParams.Page
	meta.CurrentCount = len(userList)
	meta.TotalCount = int(totalCount)
	meta.TotalPage = (int(totalCount) / queryParams.Limit) + 1
	RES_LIST_SUCCESS(c, userList, meta)
}
