package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	Success(c, 0, "success", "user info", 1)
}
