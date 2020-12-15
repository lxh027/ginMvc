package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddLevel(c *gin.Context) {
	levelModel := model.Level{}
	levelValidate := validate.LavelValidate

	if res, err := levelValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	var levelJson model.Level
	if err := c.ShouldBind(&levelJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	res := levelModel.AddLevel(levelJson)
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetAllLevel(c *gin.Context) {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
	}{}
	levelModel := model.Level{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := levelModel.GetAllLevel(format.Offset, format.Limit)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}