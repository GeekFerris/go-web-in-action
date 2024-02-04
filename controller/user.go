package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-in-action/pkg/logger"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	logger.Write("日志信息", "user")
	num1, num2 := 1, 0
	num3 := num1 / num2
	Success(c, 0, "success", num3, 1)
}
