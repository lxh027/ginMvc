package model

import "OnlineJudge/app/common"

type CarModelInfo struct {
	Mno 		int `json:"mno" form:"mno"`
	Name 		string `json:"name" form:"name"`
	Model 		string `json:"model" form:"model"`
	Total 		int `json:"total" form:"total"`
	Base 		float32 `json:"base" form:"base"`
	Cost 		float32 `json:"cost" form:"cost"`
	OverdueCost	float32 `json:"overdue_cost" form:"overdue_cost"`
	Deposit 	float32 `json:"deposit" form:"deposit"`
}

func (CarModelInfo) TableName() string {
	return "model_info"
}

func (model *CarModelInfo) GetAllModelInfo(offset int, limit int, name string, carModel string) common.ReturnType {
	var carModelInfo []CarModelInfo

	where := "name like ? AND model like ?"
	var count int
	db.Model(&CarModelInfo{}).Where(where, "%"+name+"%", "%"+carModel+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%", "%"+carModel+"%").
		Find(&carModelInfo).
		Error


	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"models": carModelInfo,
				"count": count,
			},
		}
	}
}



