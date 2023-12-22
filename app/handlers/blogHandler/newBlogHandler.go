package blogHandler

import (
	"gin-new/app/models/blogModel"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
)

func NewBlogHandler(ctx *gin.Context) {
	data := types.BlogNewReq{}
	err := ctx.BindJSON(&data)
	blog := data.ToBlogModel()
	if err == nil {
		_, errSql := blogModel.NewBlogContext(&blog)
		if errSql == nil {
			ctx.JSON(200, types.CommonRps{
				Code: 200,
				Mes:  "success",
			})
			return
		} else {
			ctx.JSON(200, types.CommonRps{
				Code: 402,
				Mes:  "failed",
			})
			return
		}
	}
	ctx.JSON(200, types.CommonRps{
		Code: 400,
		Mes:  "failed",
	})
	return
}
