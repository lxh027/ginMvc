package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddModel(c *gin.Context) {
	modelModel := model.Model{}
	modelValidate := validate.ModelValidate

	if res, err := modelValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	var modelJson model.Model
	if err := c.ShouldBind(&modelJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	res := modelModel.AddModel(modelJson)
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func EditModelLevel(c *gin.Context) {
	modelModel := model.Model{}
	modelValidate := validate.ModelValidate

	if res, err := modelValidate.Validate(c, "edit"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}
	format := struct {
		Lno int `json:"lno" form:"lno"`
		Mno int `json:"mno" form:"mno"`
	}{}

	if c.ShouldBind(&format) == nil {
		res := modelModel.EditModelLevel(format.Mno, format.Lno)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, format))
	}
}

func GetAllModel(c *gin.Context)  {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
		Where struct{
			Name string `json:"name"`
			Model string `json:"model"`
		} `json:"where"`
	}{}
	modelModel := model.CarModelInfo{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := modelModel.GetAllModelInfo(format.Offset, format.Limit, format.Where.Name, format.Where.Model)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}
