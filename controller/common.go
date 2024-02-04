package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func Success(c *gin.Context, code int, msg string, data interface{}, count int64) {
	json := &Res{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(http.StatusOK, json)
}

func Error(c *gin.Context, code int, msg string) {
	json := &Res{
		Code: code,
		Msg:  msg,
	}
	c.JSON(http.StatusInternalServerError, json)
}
