package router

import (
	"github.com/gin-gonic/gin"
	"go-web-in-action/controller"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/info", controller.UserController{}.GetUserInfo)
		user.GET("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, "")
		})

		user.POST("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, "")
		})

		user.PUT("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, "")
		})

		user.DELETE("/", func(context *gin.Context) {
			context.JSON(http.StatusOK, "")
		})
	}

	return router
}
