package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddStaff(c *gin.Context) {
	staffModel := model.Staff{}
	staffValidate := validate.StaffValidate

	if res, err := staffValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	var staffJson model.Staff
	if err := c.ShouldBind(&staffJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	res := staffModel.AddStaff(staffJson)
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func FindAllStaff(c *gin.Context)  {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
		Where struct{
			Name string `json:"name"`
			Licence string `json:"licence"`
			Number string `json:"number"`
		} `json:"where"`
	}{}
	staffModel := model.Staff{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := staffModel.FindAllStaff(format.Offset, format.Limit, format.Where.Name, format.Where.Licence, format.Where.Number)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}

func GetStaffByID(c *gin.Context)  {
	staffModel := model.Staff{}
	staffValidate := validate.StaffValidate

	if res, err := staffValidate.Validate(c, "findByID"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	customerID, _ := strconv.ParseUint(c.PostForm("customer_id"), 10, 64)
	res := staffModel.GetStaffByID(uint(customerID))

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetStaffByLicence(c *gin.Context)  {
	staffModel := model.Staff{}
	staffValidate := validate.StaffValidate

	if res, err := staffValidate.Validate(c, "findByLicence"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	licence := c.PostForm("licence")
	res := staffModel.GetStaffByLicence(licence)

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetStaffByNumber(c *gin.Context)  {
	staffModel := model.Staff{}
	staffValidate := validate.StaffValidate

	if res, err := staffValidate.Validate(c, "findByNumber"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	number := c.PostForm("number")
	res := staffModel.GetStaffByNumber(number)

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}