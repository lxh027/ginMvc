package model

import (
	"OnlineJudge/app/common"
	"time"
)

type Taken struct {
	ID 			int `json:"id" form:"id"`
	Cno 		int `json:"cno" form:"cno"`
	Sno			int	`json:"sno" form:"sno"`
	TakenPos	int `json:"taken_pos" form:"taken_pos"`
	ReturnPos	int `json:"return_pos" form:"return_pos"`
	TakenTime	time.Time `json:"taken_time" form:"taken_time"`
	ReturnTime 	time.Time `json:"return_time" form:"return_time"`
	Status 		int `json:"status" form:"status"`
}

func (model *Taken) GetAllTaken(offset int, limit int) common.ReturnType {
	var takens []Taken

	var count int
	db.Model(&Taken{}).Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Find(&takens).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"takens": takens,
				"count": count,
			},
		}
	}
}

func (model *Taken) AddTaken(data Taken)common.ReturnType  {
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}

func (model *Taken) UpdateTaken(id int, data map[string]interface{}) common.ReturnType {
	err := db.Model(&Taken{}).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "更新失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "更新成功", Data: 1}
	}
}