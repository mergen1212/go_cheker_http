package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getHostPort() (string, int) {
	return "127.0.0.1", 8080
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:key", sayAlliveView)
	return router
}

func sayAlliveView(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{})
}
func maim() {
	host, port := getHostPort()
	router := getRouter()
	router.Run(fmt.Sprintf("%s %d", host, port))

}
