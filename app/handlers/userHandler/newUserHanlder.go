package userHandler

import (
	"gin-new/app/models"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
)

func NewUser(ctx *gin.Context) {
	u := types.UserNewReq{}
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(200, models.CommonRps{
			Code: 401,
			Mes:  "failed",
		})
		return
	}
	newUser := models.User{Username: u.Uname, UserPWD: u.Pwd}
	resutl := models.TX.Create(newUser).RowsAffected
	if resutl > 0 {
		ctx.JSON(200, models.CommonRps{
			Code: 200,
			Mes:  "success",
		})
		return
	} else {
		ctx.JSON(200, models.CommonRps{
			Code: 401,
			Mes:  "failed",
		})
		return
	}
}
