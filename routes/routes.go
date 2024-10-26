package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", func(c *gin.Context) {

		const result = "Hello, World!"

		c.JSON(200, gin.H{
			"message": result,
		})
	})
}

func CreateRoutes(r *gin.Engine) {
	r.POST("/create", func(c *gin.Context) {

		var requestBody struct {
			Name string `json:"name"`
		}

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"name":    requestBody.Name,
		})

	})
}
