package userHandler

import (
	"gin-new/app/models"
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
	"log"
)

func FindHandler(ctx *gin.Context) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	data := types.UserLoginReq{}
	err := ctx.BindJSON(&data)
	if err != nil {
		log.Println("[Routers]\033[1;46;0m" + err.Error() + "\033[0m")
		ctx.JSON(200, models.CommonRps{
			Code: 402,
			Mes:  "failed",
			Data: nil,
		})
		return
	}
	userData := models.FindOne(data.Uname)
	if userData.Id < 1 {
		ctx.JSON(200, models.CommonRps{
			Code: 401,
			Mes:  "failed",
			Data: nil,
		})
		return
	}
	log.Println(userData.ToString())
	ctx.JSON(200, models.CommonRps{
		Code: 200,
		Mes:  "success",
		Data: userData,
	})
}
