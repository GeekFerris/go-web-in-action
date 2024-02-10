package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-in-action/model"
	"strconv"
)

type FoodController struct{}

func (f FoodController) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	model.GetById(id)
}

func (f FoodController) Create(c *gin.Context) {

}

func (f FoodController) Update(c *gin.Context) {

}

func (f FoodController) Delete(c *gin.Context) {

}
