package routes

import (
	"github.com/Tijanieneye10/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(server *gin.Engine) {

	//Group a router
	authenticatedGroup := server.Group("/")
	//Set middleware for the group
	authenticatedGroup.Use(middlewares.AuthMiddleware)

	//add routes to group
	authenticatedGroup.DELETE("/events/:id", deleteEvent)
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events/:id", updateEvent)

	server.GET("/events", homePage)
	server.GET("/events/:id", getEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/events", middlewares.AuthMiddleware, createEvent)
	//server.POST("/events", middlewares.AuthMiddleware, createEvent)
	server.PUT("/events/:id", updateEvent)

	//Users
	server.POST("/signup", registerUser)
	server.POST("/login", login)
}
