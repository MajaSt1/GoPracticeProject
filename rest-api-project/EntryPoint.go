package restapiproject

import (
	"net/http"

	"example.com/note/rest-api-project/db"
	"example.com/note/rest-api-project/models"
	"github.com/gin-gonic/gin"
)

func ShowRestApiExample() {
	db.InitDB()
	server := gin.Default() // setup preconfigured http server engine

	server.GET("/events", getEvents) // GET, POST, PUT, PATCH, DELETE - function will be invoked by server.GET
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost:8080
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event) //store the data in event variable - to make available client must send the same structure of event model
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
