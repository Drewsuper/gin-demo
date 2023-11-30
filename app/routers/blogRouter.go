package routers

import "github.com/gin-gonic/gin"

func init_blog_router(g *gin.RouterGroup) {
	g.POST("/blog/new")
	g.POST("/blog/delete")
	g.POST("/blog/update")
}
