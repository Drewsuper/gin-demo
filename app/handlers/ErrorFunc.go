package handlers

import (
	"gin-new/app/types"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func MyRecover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[System]: %v", r)
			debug.PrintStack()
			ctx.JSON(200, types.CommonRps{
				Code: 400,
				Mes:  "system hava some error",
			})
			ctx.Abort()
		}
	}()
	ctx.Next()
}
