package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddGarage(c *gin.Context) {
	garageModel := model.Garage{}
	garageValidate := validate.GarageValidate

	if res, err := garageValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	var garageJson model.Garage
	if err := c.ShouldBind(&garageJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	res := garageModel.AddGarage(garageJson)
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetAllGarage(c *gin.Context) {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
		Where struct{
			Position string `json:"position"`
			Director string `json:"director"`
		} `json:"where"`
	}{}
	garageModel := model.GarageInfo{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := garageModel.GetAllGarageInfo(format.Offset, format.Limit, format.Where.Position, format.Where.Director)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}

func GetNoFullGarage(c *gin.Context) {
	garageModel := model.GarageInfo{}
	res := garageModel.GetNoFullGarageInfo()
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
}

