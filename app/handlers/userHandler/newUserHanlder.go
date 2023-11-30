package userHandler

import (
	"gin-new/app/models"
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
	newUser := models.User{Username: u.Uname, UserPWD: u.Pwd}
	log.Printf("%v", newUser)
	result, _ := models.InsertUser(&newUser)
	if result > 0 {
		ctx.JSON(200, types.CommonRps{
			Code: 200,
			Mes:  "success",
		})
		return
	} else {
		ctx.JSON(200, types.CommonRps{
			Code: 401,
			Mes:  "failed",
		})
		return
	}
}
