package model

import "go-web-in-action/dao"

// Food 字段注释说明了gorm库把struct字段转换为表字段名长什么样子。
type Food struct {
	Model
	Name   string  //表字段名为：name
	Price  float64 //表字段名为：price
	TypeId int     //表字段名为：type_id
}

// TableName 设置表名，可以通过给Food struct类型定义 TableName函数，返回一个字符串作为表名
func (f Food) TableName() string {
	return "foods"
}

func GetById(id int) Food {
	db := dao.GetDB()
	food := Food{}
	db.Where("id = ?", id).First(&food)
	return food
}
