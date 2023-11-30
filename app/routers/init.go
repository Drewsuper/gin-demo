package routers

import "github.com/gin-gonic/gin"

func InitRouters(r *gin.Engine) {
	group := r.Group("/v1/api")
	init_user_routers(group)
	//init_blog_router(group)
}
