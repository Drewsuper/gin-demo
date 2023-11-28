package middleware

import (
	"gin-new/app/types"
	"gin-new/app/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Abort()
			ctx.JSON(400, types.CommonRps{
				Code: 400,
				Mes:  "no auth request",
			})
			return
		}
		_, err := utils.ParseToken(auth)
		if err != nil {
			ctx.Abort()
			ctx.JSON(400, types.CommonRps{
				Code: 400,
				Mes:  "no auth request",
			})
			return
		}
		ctx.Next()
	}
}
