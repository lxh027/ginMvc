package model

import "OnlineJudge/app/common"

type Level struct {
	Lno 		int `json:"lno" form:"lno"`
	Base 		float32 `json:"base" form:"base"`
	Cost 		float32 `json:"cost" form:"cost"`
	OverdueCost	float32 `json:"overdue_cost" form:"overdue_cost"`
	Deposit 	float32 `json:"deposit" form:"deposit"`
}

func (model *Level) GetAllLevel(offset int, limit int) common.ReturnType {
	var levels []Level

	var count int
	db.Model(&Level{}).Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Find(&levels).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"levels": levels,
				"count": count,
			},
		}
	}
}


func (model *Level) AddLevel(data Level)common.ReturnType  {
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

