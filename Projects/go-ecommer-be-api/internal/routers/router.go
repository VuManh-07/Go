package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1/api")
	{
		// Define v1 routes here
		v1.GET("/ping", ping)
	}

	v2 := r.Group("v2/api")
	{
		// Define v2 routes here
		v2.GET("/pong", pong)
	}
	return r
}

func ping(c *gin.Context) {
	// Return JSON response
	uuid := c.Query("uuid")
	name := c.DefaultQuery("name", "Manhhhh")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping..ping",
		"uuid":    uuid,
		"name":    name,
	})
}

func pong(c *gin.Context) {
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong..pong",
	})
}
