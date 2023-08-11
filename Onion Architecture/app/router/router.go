package router

import (
	"OnionPractice/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	todo := router.Group("/todos")
	{
		todo.GET("/:id", controller.Get)
		todo.POST("", controller.Create)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	return router
}
