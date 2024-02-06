package model

import "go-web-in-action/dao"

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
// 在这里User类型可以代表mysql users表
type User struct {
	//字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是golang的基础语法，不同的库会定义不同的标签，有不同的含义
	Id       int    `gorm:"column:ID; PRIMARY_KEY"`
	Name     string //`gorm:"column:NAME"`
	Username string //`gorm:"column:USERNAME"`
	Mobile   string //`gorm:"column:MOBILE"`
	// TypeId     int  表字段名为：type_id
}

// TableName 设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (User) TableName() string {
	return "s_user"
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.GetDB().Where("id = ?", id).First(&user).Error
	return user, err
}
