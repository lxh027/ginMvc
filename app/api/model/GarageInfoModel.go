package model

import "OnlineJudge/app/common"

type GarageInfo struct {
	Gno      int    `json:"gno" form:"gno"`
	Position string `json:"position" form:"position"`
	Director string `json:"director" form:"director"`
	Number   string `json:"number" form:"number"`
	Capacity int    `json:"capacity" form:"capacity"`
	Volume   int    `json:"volume" form:"volume"`
}

func (model *GarageInfo) GetNoFullGarageInfo() common.ReturnType {
	var garageInfo []GarageInfo

	where := "capacity > volume"
	var count int
	db.Model(&GarageInfo{}).Where(where).Count(&count)

	err := db.
		Where(where).
		Find(&garageInfo).
		Error


	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"garages": garageInfo,
				"count": count,
			},
		}
	}
}

func (model *GarageInfo) GetAllGarageInfo(offset int, limit int, position string, director string) common.ReturnType {
	var garageInfo []GarageInfo

	where := "position like ? AND director like ?"
	var count int
	db.Model(&GarageInfo{}).Where(where, "%"+position+"%", "%"+director+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+position+"%", "%"+director+"%").
		Find(&garageInfo).
		Error


	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"garages": garageInfo,
				"count": count,
			},
		}
	}
}