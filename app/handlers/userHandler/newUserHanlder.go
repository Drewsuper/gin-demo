package userHandler

import (
	"gin-new/app/models/userModel"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
	"log"
)

func NewUser(ctx *gin.Context) {
	u := types.UserNewReq{}
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.JSON(200, types.CommonRps{
			Code: 401,
			Mes:  "failed",
		})
		return
	}
	newUser := userModel.User{Username: u.Uname, UserPWD: u.Pwd}
	log.Printf("%v", newUser)
	_, errOne := userModel.InsertUser(&newUser)
	if errOne == nil {
		ctx.JSON(200, types.CommonRps{
			Code: 200,
			Mes:  "success",
		})
		return
	}
	ctx.JSON(200, types.CommonRps{
		Code: 401,
		Mes:  "failed",
	})
	return
}
