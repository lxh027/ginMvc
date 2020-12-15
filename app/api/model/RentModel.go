package model

import (
	"OnlineJudge/app/common"
	"time"
)

type Rent struct {
	ID 			int `json:"id" form:"id"`
	Cno 		int `json:"cno" form:"cno"`
	CustomerID	int	`json:"customer_id" form:"customer_id"`
	Lno 		int `json:"lno" form:"lno"`
	RentPos		int `json:"rent_pos" form:"rent_pos"`
	ReturnPos	int `json:"return_pos" form:"return_pos"`
	RentTime	time.Time `json:"rent_time" form:"rent_time"`
	ExpectTime	time.Time `json:"expect_time" form:"expect_time"`
	ReturnTime 	time.Time `json:"return_time" form:"return_time"`
	Pay 		float32 `json:"pay" form:"pay"`
	Status 		int `json:"status" form:"status"`
}

func (model *Rent) GetAllRent(offset int, limit int) common.ReturnType {
	var rents []Rent

	var count int
	db.Model(&Rent{}).Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Find(&rents).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"rents": rents,
				"count": count,
			},
		}
	}
}

func (model *Rent) AddRent(data Rent)common.ReturnType  {
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

func (model *Rent) UpdateRent(id int, data map[string]interface{}) common.ReturnType {
	err := db.Model(&Rent{}).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "更新失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "更新成功", Data: 1}
	}
}