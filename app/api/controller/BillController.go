package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllRent(c *gin.Context)  {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
	}{}
	rentModel := model.Rent{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := rentModel.GetAllRent(format.Offset, format.Limit)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}

func GetAllTaken(c *gin.Context)  {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
	}{}
	takenModel := model.Taken{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := takenModel.GetAllTaken(format.Offset, format.Limit)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}


