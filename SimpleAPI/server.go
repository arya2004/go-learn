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

	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)

)


func setupLogOutput(){
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}


func main() {
	setupLogOutput()
	server := gin.New()


	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	// server.GET("/test", func(ctx * gin.Context){
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "OK!!",
	// 	})
	// })
	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})


	apiROutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiROutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})


		apiROutes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"} )
			}
			
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hola World")
	})


	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}
	

	server.Run("0.0.0.0:3333")

}