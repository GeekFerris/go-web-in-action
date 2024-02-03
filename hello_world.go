package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/index", func(context *gin.Context) {
		context.String(200, "Hello, World")
	})

	err := router.Run(":8080")
	if err != nil {
	}

}
