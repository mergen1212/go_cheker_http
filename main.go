package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_cheker_http/pkg/db"
	"net/http"
)

func getHostPort() (string, int) {
	return "127.0.0.1", 8080
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/:key", sayAliveView)
	return router
}

func sayAliveView(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{})
}
func main() {
	host, port := getHostPort()
	router := getRouter()
	db, err := db.GetDB()
	if err != nil {
		panic(err)
	}
	err = db.PrepareDB()
	if err != nil {
		panic(err)
	}
	router.Run(fmt.Sprintf("%s %d", host, port))
}
