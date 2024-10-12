package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/internal/routes"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	fmt.Println("Server is starting on port 8080")
	r.Run(":8080")
}
