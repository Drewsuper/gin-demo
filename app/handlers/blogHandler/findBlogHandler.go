package blogHandler

import (
	"gin-new/app/models/blogModel"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
)

func FindBlogHandler(ctx *gin.Context) {
	data := types.BlogFindReq{}
	err := ctx.BindJSON(&data)
	if err == nil {
		blog, err := blogModel.FindOneById(data.ID)
		if err == nil {
			ctx.JSON(200, types.CommonRps{
				Code: 200,
				Mes:  "success",
				Data: blog,
			})
			return
		}
	}
	ctx.JSON(200, types.CommonRps{
		Code: 200,
		Mes:  "failed",
	})
	return
}
