package model

import "OnlineJudge/app/common"

type Staff struct {
	Sno 		int 	`json:"sno" form:"sno"`
	Name 		string 	`json:"name" form:"name"`
	Sex 		int 	`json:"sex" form:"sex"`
	Number		string	`json:"number" form:"number"`
	Licence 	string  `json:"licence" form:"licence"`
}

// 添加员工
func (model *Staff) AddStaff(data Staff)common.ReturnType  {
	staff := Staff{}
	// 判断身份证号和电话号码是否已存在
	if err := db.Where("licence = ? OR number = ?", data.Licence, data.Number).First(&staff).Error; err == nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "身份证或号码已存在",  Data: staff}
	}
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

func (model *Staff) FindAllStaff(offset int, limit int, name string, licence string, number string) common.ReturnType  {
	var staffs []Staff
	where := "name like ? AND licence like ? AND number like ?"

	var count int
	db.Model(&Staff{}).Where(where, "%"+name+"%", "%"+licence+"%", "%"+number+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%", "%"+licence+"%", "%"+number+"%").
		Find(&staffs).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"staffs": staffs,
				"count": count,
			},
		}
	}}

// 通过ID获取用户
func (model *Staff) GetStaffByID(staffID uint)common.ReturnType  {
	staff := Staff{}
	err := db.Where("sno = ?", staffID).First(&staff).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: staff}
	}
}

// 通过身份证号获取用户
func (model *Staff) GetStaffByLicence(licence string)common.ReturnType  {
	staff := Staff{}
	err := db.Where("licence = ?", licence).First(&staff).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: staff}
	}
}

// 通过电话获取用户
func (model *Staff) GetStaffByNumber(number string)common.ReturnType  {
	staff := Staff{}
	err := db.Where("number = ?", number).First(&staff).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功", Data: staff}
	}
}
