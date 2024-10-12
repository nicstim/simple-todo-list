package routes

import (
	"github.com/gin-gonic/gin"
	"main/internal/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("", handlers.GetAllTask)
	r.GET("/:id", handlers.GetTaskById)
	r.POST("", handlers.AddTask)
	r.PUT("/:id", handlers.ChangeTaskStatus)
}
