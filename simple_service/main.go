package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
