package main

import (
	"gin-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//fmt.Println("Starting the server on http://localhost:8080")

	r := gin.Default()

	routes.RegisterRoutes(r)
	routes.CreateRoutes(r)

	r.Run()
}
