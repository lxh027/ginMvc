package controller

import (
	"OnlineJudge/app/api/model"
	"OnlineJudge/app/api/validate"
	"OnlineJudge/app/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCustomer(c *gin.Context) {
	customerModel := model.Customer{}
	customerValidate := validate.CustomerValidate

	if res, err := customerValidate.Validate(c, "add"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	var customerJson model.Customer
	if err := c.ShouldBind(&customerJson); err != nil {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "数据绑定模型错误", err.Error()))
		return
	}

	res := customerModel.AddCustomer(customerJson)
	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func FindAllCustomer(c *gin.Context)  {
	format := struct {
		Offset int `json:"offset" form:"offset"`
		Limit int `json:"limit" form:"limit"`
		Where struct{
			Name string `json:"name"`
			Licence string `json:"licence"`
			Number string `json:"number"`
		} `json:"where"`
	}{}
	customerModel := model.Customer{}
	if c.ShouldBind(&format) == nil {
		format.Offset = (format.Offset-1)*format.Limit
		res := customerModel.FindAllCustomer(format.Offset, format.Limit, format.Where.Name, format.Where.Licence, format.Where.Number)
		c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	}
	return
}

func GetCustomerByID(c *gin.Context)   {
	customerModel := model.Customer{}
	customerValidate := validate.CustomerValidate

	if res, err := customerValidate.Validate(c, "findByID"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	customerID, _ := strconv.ParseUint(c.PostForm("customer_id"), 10, 64)
	res := customerModel.GetCustomerByID(uint(customerID))

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetCustomerByLicence(c *gin.Context)  {
	customerModel := model.Customer{}
	customerValidate := validate.CustomerValidate

	if res, err := customerValidate.Validate(c, "findByLicence"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	licence := c.PostForm("licence")
	res := customerModel.GetCustomerByLicence(licence)

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}

func GetCustomerByNumber(c *gin.Context)  {
	customerModel := model.Customer{}
	customerValidate := validate.CustomerValidate

	if res, err := customerValidate.Validate(c, "findByNumber"); !res {
		c.JSON(http.StatusOK, common.ApiReturn(common.CODE_ERROE, "输入信息不完整或有误", err.Error()))
		return
	}

	number := c.PostForm("number")
	res := customerModel.GetCustomerByLicence(number)

	c.JSON(http.StatusOK, common.ApiReturn(res.Status, res.Msg, res.Data))
	return
}