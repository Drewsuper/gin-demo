package main

import (
	"gin-new/app/config"
	"gin-new/app/models"
	"gin-new/app/routers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Printf("\033[0;36m[System]: running...\033[0m")
	errC := config.InitConfig()
	if errC != nil {
		os.Exit(-2)
	}
	errB := models.InitDataBase()
	if errB != nil {
		os.Exit(-1)
	}
	log.Println("\u001B[0;33m[System]: routers  is initializing ...\033[0m")
	time.Sleep(100)
	r := gin.Default()
	routers.InitRouters(r)
	err := r.Run(config.Config.Host + ":" + config.Config.Port)
	if err != nil {
		log.Print(err)
		return
	}
}
