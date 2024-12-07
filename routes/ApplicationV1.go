package routes

import (
	"trinity-app/controllers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplicationV1Router(router *gin.Engine) {
	// add swagger
	url := ginSwagger.URL("http://localhost:5001/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	router.GET("/", controllers.Healthcheck)
	api := router.Group("/api/v1")
	{	
		
		auth := api.Group("/trinity")
		{
			///////////////// TRINITY API
			auth.GET("/create-table", controllers.MigrateModel)
			auth.POST("/member-register", controllers.MemberRegister)
			auth.POST("/silver-register", controllers.SilverRegister)
			auth.POST("/create-coupon", controllers.CreateCoupon)
			auth.POST("/member-login", controllers.MemberLogin)
			auth.GET("/coupon-list", controllers.GetCouponList)
			///////////////// TRINITY API


			// auth.POST("/member-login-social", controllers.MemberLoginSocial)
			// auth.POST("/admin-login", controllers.AdminLogin)
			// auth.POST("/member-register-social", controllers.MemberRegisterSocial)
		}
		profile := api.Group("/profile")
		{
			profile.GET("/get-profile-list", controllers.GetCouponList)
			profile.GET("/get-profile", controllers.GetUserProfile)
			profile.GET("/get-profile-by/:id", controllers.GetUserProfileById)
			profile.PUT("/update-profile", controllers.UpdateProfile)
			profile.PUT("/update-fcm-token", controllers.UpdateFcmToken)
		}
		socialInfo := api.Group("/social-info")
		{
			socialInfo.POST("/create-or-update", controllers.CreateOrUpdateSocialInfo)
			socialInfo.GET("/get-social-info", controllers.GetSocialInfo)
		}
	}
}
