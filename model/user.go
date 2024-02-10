package model

import "go-web-in-action/dao"

// User 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
// 在这里User类型可以代表mysql users表
type User struct {
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//字段定义后面使用两个反引号``包裹起来的字符串部分叫做标签定义，这个是golang的基础语法，不同的库会定义不同的标签，有不同的含义
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	//这里golang定义的Username变量和MYSQL表字段username一样，他们的名字可以不一样。
	Model
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	//CreateTime int64 `gorm:"column:createtime"`
}

// TableName 设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (User) TableName() string {
	return "users"
}

func GetUserTest(id int) User {
	//var user User
	//err := dao.GetDB().Where("id = ?", id).First(&user).Error
	//return user, err
	db := dao.GetDB()
	u := User{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'sysadmin') LIMIT 1
	//dao.GetDB().Where("username = ?", "sysadmin").First(&u)
	db.Take(&u)

	return u
}

func createUser(username, password string) (uint, error) {
	user := User{
		Username: username,
		Password: password,
	}
	err := dao.GetDB().Create(&user).Error
	return user.ID, err
}
