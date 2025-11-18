package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Backend service is running...")

	_ = godotenv.Load()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to workflow backend. 001",
		})
	})
	r.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "workflow backend /api",
		})
	})

	port := os.Getenv("BACKEND_PORT")
	r.Run(":" + port)
}
