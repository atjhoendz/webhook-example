package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "grade": value})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		}
	})

	r.POST("/webhook", func(c *gin.Context) {
		var payload struct {
			Name  string `json:"name" binding:"required"`
			Grade string `json:"grade" binding:"required"`
		}

		if c.Bind(&payload) == nil {
			db[payload.Name] = payload.Grade
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
