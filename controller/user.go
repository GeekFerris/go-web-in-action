package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	//idStr := c.Param("id")
	//id, _ := strconv.Atoi(idStr)
	//user := model.GetUserTest(id)
	//Success(c, 0, "success", user, 1)
}
