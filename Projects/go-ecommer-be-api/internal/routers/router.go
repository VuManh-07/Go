package routers

import (
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	c := controller.NewUserController()

	v1 := r.Group("v1/api")
	{
		// Define v1 routes here
		// v1.GET("/ping", ping)
		v1.GET("/user", c.GetUserName)
	}

	// v2 := r.Group("v2/api")
	// {
	// 	// Define v2 routes here
	// 	v2.GET("/pong", pong)
	// }
	return r
}
