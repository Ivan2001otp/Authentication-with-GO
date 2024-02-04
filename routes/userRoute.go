package routes

import (
	"github.com/Ivan2001otp/Authentication-with-GO/controllers"
	"github.com/Ivan2001otp/Authentication-with-GO/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/user_id", controllers.GetUser())
}
