package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Gateway service is running...")

	_ = godotenv.Load()

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome to workflow gateway.",
		})
	})
	r.GET("/v1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "workflow gateway /v1",
		})
	})

	port := os.Getenv("GATEWAY_PORT")
	r.Run(":" + port)
}
