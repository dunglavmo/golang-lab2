package routes

import (
	"your-module-name/controllers"
	"your-module-name/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.LogMiddleware)

	r.POST("/login", controllers.Authenticate)
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PUT("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	return r
}
