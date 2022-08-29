package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users")
	router.POST("/users/newUser")
	// router.GET("/question")
	// router.POST("/question")
	router.Run("localhost:3000")

}
