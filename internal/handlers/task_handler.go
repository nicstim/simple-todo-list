package handlers

import (
	"github.com/gin-gonic/gin"
	"main/internal/models"
	"net/http"
)

var tasks = []models.Task{
	{ID: "1", Title: "Написать свой первый сервис на Go", Status: true},
	{ID: "2", Title: "Собраться в поездку", Status: false},
}

func GetAllTask(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tasks)
}

func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	for _, taskObject := range tasks {
		if taskObject.ID == id {
			c.IndentedJSON(http.StatusOK, taskObject)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func ChangeTaskStatus(c *gin.Context) {
	id := c.Param("id")
	for index, taskObject := range tasks {
		if taskObject.ID == id {
			tasks[index].Status = true
			c.IndentedJSON(http.StatusOK, tasks[index])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}
