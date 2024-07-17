package restapiproject

import (
	"example.com/note/rest-api-project/db"
	"example.com/note/rest-api-project/routes"
	"github.com/gin-gonic/gin"
)

func ShowRestApiExample() {
	db.InitDB()
	server := gin.Default() // setup preconfigured http server engine
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}