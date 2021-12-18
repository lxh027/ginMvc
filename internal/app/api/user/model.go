package user

import (
	"mvc/internal/constants"
	"mvc/internal/global"
	"mvc/tools/formatter"
)

type Dao struct{}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

func (dao Dao) AddUser(user User) formatter.ReturnType {
	// 判断昵称是否已存在
	err := global.MysqlClient.
		Where("username = ?", user.Username).
		First(&User{}).
		Error
	if err == nil {
		return formatter.ReturnType{Status: constants.CodeError, Msg: "昵称已存在", Data: user}
	}
	// 创建记录
	err = global.MysqlClient.Create(&user).Error
	if err != nil {
		return formatter.ReturnType{Status: constants.CodeError, Msg: "创建失败", Data: err.Error()}
	}
	return formatter.ReturnType{Status: constants.CodeSuccess, Msg: "创建成功", Data: nil}
}

func (dao Dao) CheckUserLogin(username string, password string) formatter.ReturnType {
	fields := []string{"id", "username"}

	var user User
	err := global.MysqlClient.
		Select(fields).
		Where("username = ? AND password = ?", username, password).
		First(&user).
		Error

	if err != nil {
		return formatter.ReturnType{Status: constants.CodeError, Msg: "用户名或密码错误", Data: nil}
	}
	return formatter.ReturnType{Status: constants.CodeSuccess, Msg: "登录成功", Data: nil}
}
