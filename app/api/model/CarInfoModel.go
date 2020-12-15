package model

import "OnlineJudge/app/common"

type CarInfo struct {
	Cno 		int `json:"cno" form:"cno"`
	Model 		string `json:"model" form:"model"`
	Licence 	string `json:"licence" form:"licence"`
	Position 	string `json:"position" form:"position"`
	Status 		int `json:"status" form:"status"`
}

func (model *CarInfo) GetAllCarInfo(offset int, limit int, licence string, position string, carModel string) common.ReturnType {
	var carInfo []CarInfo

	where := "licence like ? AND position like ? AND model like ?"
	var count int
	db.Model(&CarInfo{}).Where(where, "%"+licence+"%", "%"+position+"%", "%"+carModel+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+licence+"%", "%"+position+"%", "%"+carModel+"%").
		Find(&carInfo).
		Error


	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"cars": carInfo,
				"count": count,
			},
		}
	}
}



