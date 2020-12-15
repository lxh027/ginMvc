package model

import "OnlineJudge/app/common"

type Model struct {
	Mno 	int 	`json:"mno" form:"mno"`
	Lno		int 	`json:"lno" form:"lno"`
	Name 	string 	`json:"name" form:"name"`
	Model 	string	`json:"model" form:"model"`
}

func (model *Model) EditModelLevel(mno int, lno int) common.ReturnType  {
	err := db.Model(&Model{}).Where("mno = ?", mno).Update("lno", lno).Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "更新失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "更新成功", Data: 0}
	}
}

func (model *Model) GetAllModel(offset int, limit int, name string, carModel string) common.ReturnType {
	var models []Model

	where := "name like ? AND model like ?"
	var count int
	db.Model(&Model{}).Where(where, "%"+name+"%", "%"+carModel+"%").Count(&count)

	err := db.Offset(offset).
		Limit(limit).
		Where(where, "%"+name+"%", "%"+carModel+"%").
		Find(&models).
		Error

	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "查询失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "查询成功",
			Data: map[string]interface{}{
				"models": models,
				"count": count,
			},
		}
	}
}

// 添加车型
func (model *Model) AddModel(data Model)common.ReturnType  {
	carModel := Model{}

	if err := db.Where("name = ? AND model = ?", data.Name, data.Model).First(&carModel).Error; err == nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "车型已存在",  Data: carModel}
	}
	// 创建记录
	err := db.Create(&data).Error
	if err != nil {
		return common.ReturnType{Status: common.CODE_ERROE, Msg: "创建失败", Data: err.Error()}
	} else {
		return common.ReturnType{Status: common.CODE_SUCCESS, Msg: "创建成功", Data: 1}
	}
}



