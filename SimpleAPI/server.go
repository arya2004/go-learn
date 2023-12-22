package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"

	"github.com/arya2004/GoLearn/SimpleAPI/controller"
	"github.com/arya2004/GoLearn/SimpleAPI/middlewares"
	"github.com/arya2004/GoLearn/SimpleAPI/service"
)

var (
	videoService service.VideoService = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)


func setupLogOutput(){
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	// server.GET("/test", func(ctx * gin.Context){
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "OK!!",
	// 	})
	// })

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hola World")
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"} )
		}
		
	})

	server.Run(":3333")

}