package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddCar(c *gin.Context) {
	carModel := model.Car{}
	carValidate := validate.CarValidate

	if res, err := carValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	var carJson model.Car
	if err := c.ShouldBind(&carJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	res := carModel.AddCar(carJson)
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetAllCar(c *gin.Context)  {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
		Where struct{
			Licence string `json:"licence"`
			Position string `json:"position"`
			Model string `json:"model"`
		} `json:"where"`
	}{}
	carModel := model.CarInfo{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := carModel.GetAllCarInfo(format.Offset, format.Limit, format.Where.Licence, format.Where.Position, format.Where.Model)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}

