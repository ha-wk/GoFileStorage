package main

import (
	"github.com/gin-gonic/gin"
	"wobot-file-storage/routes"
)

func main() {
	r := gin.Default()          //initializing GIN object and passing to router
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
