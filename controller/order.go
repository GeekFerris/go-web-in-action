package controller

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

type Condition struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (o OrderController) GetById(c *gin.Context) {
	id := c.Param("id")
	Success(c, 0, "success", id, 1)
}

func (o OrderController) QueryByJson(c *gin.Context) {
	//param := make(map[string]any)
	//err := c.BindJSON(&param)
	//if err == nil {
	//	Success(c, 0, "success", param, 1)
	//} else {
	//	Error(c, 0, "error")
	//}
	search := &Condition{}

	err := c.BindJSON(&search)

	if err == nil {
		Success(c, 0, "success", search, 1)
	} else {
		Error(c, 0, "error")
	}
}
