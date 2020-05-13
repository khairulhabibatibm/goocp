package routes

import (
	"example.com/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		v1.GET("master", controllers.GetTodos)
	}

	return r
}
