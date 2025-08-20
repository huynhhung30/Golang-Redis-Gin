package routes

import (
	"Golang-Redis-Gin/internal/controllers"
	"Golang-Redis-Gin/internal/di"
	"Golang-Redis-Gin/internal/redis"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)
func ApplicationV1Router(db *gorm.DB,r *gin.Engine, rdb  *redis.RedisCache) *gin.Engine  {

	// add swagger
	url := ginSwagger.URL("http://localhost:5001/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	r.GET("/", controllers.Healthcheck)
	userHandler := di.InitUser(db, rdb)
	api := r.Group("/api/v1")
	{	
		
		member := api.Group("/member")
		{
			/////////////////  API
			member.GET("/users/:id", userHandler.GetProfile)
			// member.GET("/create-table", controllers.MigrateTable)
			// member.POST("/member-register", controllers.MemberRegister)
			// auth.POST("/silver-register", controllers.SilverRegister)
			// auth.POST("/create-coupon", controllers.CreateCoupon)
			// auth.POST("/member-login", controllers.MemberLogin)
			// auth.GET("/coupon-list", controllers.GetCouponList)
			///////////////// TRINITY API


			// auth.POST("/member-login-social", controllers.MemberLoginSocial)
			// auth.POST("/admin-login", controllers.AdminLogin)
			// auth.POST("/member-register-social", controllers.MemberRegisterSocial)
		}
		// profile := api.Group("/profile")
		// {
		// 	profile.GET("/get-profile-list", controllers.GetCouponList)
		// 	profile.GET("/get-profile", controllers.GetUserProfile)
		// 	profile.GET("/get-profile-by/:id", controllers.GetUserProfileById)
		// 	profile.PUT("/update-profile", controllers.UpdateProfile)
		// 	profile.PUT("/update-fcm-token", controllers.UpdateFcmToken)
		// }
		// socialInfo := api.Group("/social-info")
		// {
		// 	socialInfo.POST("/create-or-update", controllers.CreateOrUpdateSocialInfo)
		// 	socialInfo.GET("/get-social-info", controllers.GetSocialInfo)
		// }
	}
	return r
}
