package userHandler

import (
	"gin-new/app/utils"
	"gin-new/app/models"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
)

func UserLoginHandler(ctx *gin.Context) {
	data := types.UserLoginReq{}
	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.JSON(200, types.CommonRps{
			Code: 200,
			Mes:  "not get params",
		})
		return
	}
	user := models.FindOne(data.Uname)
	if user.Id == 0 {
		ctx.JSON(200, types.CommonRps{
			Code: 400,
			Mes:  "login failed",
		})
		return
	}
	if user.UserPWD == data.PWD {
		token_string, err := utils.NewToken(user.Id)
		if err == nil {
			ctx.JSON(200, types.CommonRps{
				Code: 200,
				Mes:  "login success",
				Data: token_string,
			})
			return
		}
		ctx.JSON(200, types.CommonRps{
			Code: 400,
			Mes:  "login failed",
		})
		return
	}
	ctx.JSON(200, types.CommonRps{
		Code: 400,
		Mes:  "login failed",
	})
	return
}
