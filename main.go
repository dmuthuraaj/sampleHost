package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	uriHome = "/home"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "welcome to home"})
}

func main() {
	r := gin.Default()
	routerGroup := r.Group("/api/v1")
	routerGroup.GET(uriHome, HomeHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}
