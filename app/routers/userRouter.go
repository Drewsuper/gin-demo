package routers

import (
	"gin-new/app/handlers/userHandler"
	"gin-new/app/middleware"
	"github.com/gin-gonic/gin"
)

func init_user_routers(e *gin.Engine) {
	e.POST("/user/login", userHandler.UserLoginHandler)
	e.POST("/user/find", middleware.Auth(), userHandler.FindHandler)
}
