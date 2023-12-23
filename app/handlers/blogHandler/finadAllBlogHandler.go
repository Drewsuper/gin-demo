package blogHandler

import "github.com/gin-gonic/gin"

func AllFindBlog (ctx *gin.Context) {
	var data []BlogModel
	err := blogModel.FindAllBlog(&data)
	if err == nil{
		ctx.JSON(200, types.CommonRps{
			Code: 200,
			Mes: "success",
			Data: data 
		})
		return
	}else {
		ctx.JSON(200,types.CommonRps{
			Code:400,
			Mes: "failed"
		})
		return
	}
}