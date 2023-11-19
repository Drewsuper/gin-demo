package routers

import (
	"gin-new/app/handlers/userHandler"
	"github.com/gin-gonic/gin"
)

func init_user_routers(e *gin.Engine) {
	e.POST("/user/new", userHandler.FindHandler)
	e.POST("/user/find", userHandler.FindHandler)
}
