package controllers

import (
	"net/http"
	"strconv"

	"your-module-name/models"
	"your-module-name/utils"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.HandleBadRequest(c, "Invalid request data")
		return
	}

	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleBadRequest(c, "Invalid todo ID")
		return
	}

	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		utils.HandleBadRequest(c, "Invalid request data")
		return
	}

	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		utils.HandleNotFound(c, "Todo not found")
		return
	}

	config.DB.Model(&todo).Updates(updatedTodo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleBadRequest(c, "Invalid todo ID")
		return
	}

	var todo models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		utils.HandleNotFound(c, "Todo not found")
		return
	}

	config.DB.Delete(&todo)
	c.Status(http.StatusNoContent)
}
