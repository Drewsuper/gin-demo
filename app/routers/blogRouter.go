package routers

import (
	"gin-new/app/handlers/blogHandler"
	"github.com/gin-gonic/gin"
)

func init_blog_router(g *gin.RouterGroup) {
	g.POST("/blog/find", blogHandler.FindBlogHandler)
	g.POST("/blog/new", blogHandler.NewBlogHandler)
	g.POST("/blog/find_all",blogHandler.AllFindBlogHandler)
}
