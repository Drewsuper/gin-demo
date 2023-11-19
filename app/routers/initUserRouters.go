package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	init_user_routers(r)
}
