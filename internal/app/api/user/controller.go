package user

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mvc/internal/constants"
	"mvc/tools/formatter"
	"mvc/tools/md5"
	"mvc/tools/validate"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var userJson User

	if err := c.ShouldBind(&userJson); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(constants.CodeError, "获取参数失败", err.Error()))
		return
	}

	userJson.Password = md5.GetMd5(userJson.Password)

	res := Dao{}.AddUser(userJson)

	c.JSON(http.StatusOK, formatter.ApiReturn(res.Status, res.Msg, res.Data))
}

func Login(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("user_id") != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(constants.CodeSuccess, "已登陆", nil))
		return
	}

	type paramType struct {
		Password string `json:"password"`
		Username string `json:"username"`
	}
	var params paramType

	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(constants.CodeError, "获取参数失败", err.Error()))
		return
	}

	// validate
	if err := Validate.ValidateMap(validate.Struct2Map(params), "login"); err != nil {
		c.JSON(http.StatusOK, formatter.ApiReturn(constants.CodeError, "参数验证失败", err.Error()))
		return
	}

	params.Password = md5.GetMd5(params.Password)
	res := Dao{}.CheckUserLogin(params.Username, params.Password)
	if res.Status == constants.CodeSuccess {
		user := res.Data.(User)
		session.Set("user_id", user.Username)
		session.Save()
	}
	c.JSON(http.StatusOK, formatter.ApiReturn(res.Status, res.Msg, nil))
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, formatter.ApiReturn(constants.CodeSuccess, "注销成功", nil))
}
