package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hi",
		})
	})

	log.Printf("Server listening on port %s", port)
	log.Fatal(r.Run(":"+port, mux))
}
