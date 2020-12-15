package model

import "OnlineJudge/app/common"

type Car struct {
	Cno		int	`json:"cno" form:"cno"`
	Gno		int `json:"gno" form:"gno"`
	Mno		int	`json:"mno" form:"mno"`
	Licence	string 	`json:"licence" form:"licence"`
	Status	int		`json:"status" form:"status"`
}

func (model *Car) GetAllCar(offset int, limit int, licence string) common.ReturnType {
	var cars []Car

	where := "licence like ?"
	var count int
	db.Model(&Car{}).Where(where, "%"+licence+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+licence+"%").
		Find(&cars).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"cars": cars,
				"count": count,
			},
		}
	}
}

// 添加车型
func (model *Car) AddCar(data Car)common.ReturnType  {
	car := Car{}

	if err := db.Where("licence = ?", data.Licence).First(&car).Error; err == nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "车牌已存在",  Data: car}
	}

	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

// 改变车辆状态
func (model *Car) UpdateCarStatus(cno int, status int) common.ReturnType {
	err := db.Model(&Car{}).Where("cno = ?", cno).Update("status", status).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "更新失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "更新成功", Data: 1}
	}
}


