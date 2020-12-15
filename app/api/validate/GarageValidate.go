package validate

import "OnlineJudge/app/common"

var GarageValidate common.Validator

func init()  {
	rules := map[string]string{
		"gno"			: "required",
		"position"		: "required",
		"number"		: "required",
		"director"		: "required",
		"capacity"		: "required",
	}

	scenes := map[string][]string {
		"add": {"position", "director", "number", "capacity"},
		"findByID"	: {"gno"},
	}
	GarageValidate.Rules = rules
	GarageValidate.Scenes = scenes
}