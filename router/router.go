package router

import (
	"github.com/gin-gonic/gin"
	"go-web-in-action/controller"
	"go-web-in-action/pkg/logger"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	router.Use(logger.Recover)

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

	order := router.Group("/order")
	{
		order.GET("/:id", controller.OrderController{}.GetById)
		order.POST("/query", controller.OrderController{}.QueryByJson)
	}

	return router
}
