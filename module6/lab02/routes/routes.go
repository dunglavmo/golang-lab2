package routes

import (
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.PUT("/todo/:idTodo", controllers.UpdateTodo)
	route.DELETE("/todo/:idTodo", controllers.DeleteTodo)
	route.GET("/todo/:idTodo", controllers.GetByID)

	route.Run()
}
