package model

import "OnlineJudge/app/common"

type Customer struct {
	CustomerID 	int 	`json:"customer_id" form:"customer_id"`
	Name 		string 	`json:"name" form:"name"`
	Sex 		int 	`json:"sex" form:"sex"`
	Number		string	`json:"number" form:"number"`
	Licence 	string  `json:"licence" form:"licence"`
}

// 添加客户
func (model *Customer) AddCustomer(data Customer)common.ReturnType  {
	customer := Customer{}
	// 判断身份证号和电话号码是否已存在
	if err := db.Where("licence = ? OR number = ?", data.Licence, data.Number).First(&customer).Error; err == nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "身份证或号码已存在",  Data: customer}
	}
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

func (model *Customer) FindAllCustomer(offset int, limit int, name string, licence string, number string) common.ReturnType  {
	var customers []Customer

	where := "name like ? AND licence like ? AND number like ?"

	var count int
	db.Model(&Customer{}).Where(where, "%"+name+"%", "%"+licence+"%", "%"+number+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%", "%"+licence+"%", "%"+number+"%").
		Find(&customers).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"customers": customers,
				"count": count,
			},
		}
	}
}

// 通过ID获取用户
func (model *Customer) GetCustomerByID(customerID uint)common.ReturnType  {
	customer := Customer{}
	err := db.Where("customer_id = ?", customerID).First(&customer).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: customer}
	}
}

// 通过身份证号获取用户
func (model *Customer) GetCustomerByLicence(licence string)common.ReturnType  {
	customer := Customer{}
	err := db.Where("licence = ?", licence).First(&customer).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: customer}
	}
}

// 通过电话获取用户
func (model *Customer) GetCustomerByNumber(number string)common.ReturnType  {
	customer := Customer{}
	err := db.Where("number = ?", number).First(&customer).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: customer}
	}
}
