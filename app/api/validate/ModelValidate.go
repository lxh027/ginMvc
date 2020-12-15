package validate

import "OnlineJudge/app/common"

var ModelValidate common.Validator

func init()  {
	rules := map[string]string{
		"mno"			: "required",
		"lno"			: "required",
		"name"			: "required",
		"model"			: "required",
	}

	scenes := map[string][]string {
		"add": {"name", "model"},
		"edit": {"mno", "lno"},
		"findByID"	: {"gno"},
	}
	ModelValidate.Rules = rules
	ModelValidate.Scenes = scenes
}