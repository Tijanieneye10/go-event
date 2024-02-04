package routes

import "github.com/gin-gonic/gin"

func RegisterRouter(server *gin.Engine) {
	server.GET("/events", homePage)
	server.GET("/events/:id", getEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)

	//Users
	server.POST("/signup", registerUser)
	server.POST("/login", login)
}
