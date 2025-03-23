package routes

import (
	"wobot-file-storage/controllers"
	"wobot-file-storage/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)           //jwt will be created here

	//routes which requires jwt auth
	auth := r.Group("/", middleware.JWTMiddleware())
	{
		auth.POST("/upload", controllers.UploadFileHandler)
		auth.GET("/storage/remaining", controllers.GetStorageHandler)
		auth.GET("/files", controllers.ListFilesHandler)
	}
}
