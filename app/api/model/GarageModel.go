package model

import "OnlineJudge/app/common"

type Garage struct {
	Gno 		int `json:"gno" form:"gno"`
	Position 	string `json:"position" form:"position"`
	Number 		string `json:"number" form:"number"`
	Director	int `json:"director" form:"director"`
	Capacity 	int `json:"capacity" form:"capacity"`
}

func (model *Garage) GetNoFullGarage() common.ReturnType {
	var garages []Garage

	var count int
	db.Model(&Garage{}).Count(&count)

	err := db.Find(&garages).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"garages": garages,
				"count": count,
			},
		}
	}
}


func (model *Garage) GetAllGarage(offset int, limit int, position string) common.ReturnType {
	var garages []Garage

	where := "position like ?"
	var count int
	db.Model(&Garage{}).Where(where, "%"+position+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+position+"%").
		Find(&garages).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"garages": garages,
				"count": count,
			},
		}
	}
}

// 添加仓库
func (model *Garage) AddGarage(data Garage)common.ReturnType  {
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}


