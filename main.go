package main

import (
	"github.com/Tijanieneye10/db"
	"github.com/Tijanieneye10/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()

	server := gin.Default()

	//Create a route an pass a the server engine
	routes.RegisterRouter(server)

	server.Run(":8080")

}
