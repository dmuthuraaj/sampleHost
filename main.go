package main

import (
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
}
