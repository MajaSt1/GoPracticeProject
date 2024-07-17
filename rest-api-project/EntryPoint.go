package restapiproject

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowRestApiExample() {
	server := gin.Default() // setup preconfigured http server engine

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE - function will be invoked by server.GET

	server.Run(":8080") // localhost:8080
}

func getEvents(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
