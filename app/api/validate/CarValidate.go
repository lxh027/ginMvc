package validate

import "OnlineJudge/app/common"

var CarValidate common.Validator

func init()  {
	rules := map[string]string{
		"cno": "required",
		"mno": "required",
		"gno": "required",
		"licence": "required",
	}

	scenes := map[string][]string {
		"add": {"mno", "licence", "gno"},
	}
	CarValidate.Rules = rules
	CarValidate.Scenes = scenes
}