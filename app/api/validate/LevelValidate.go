package validate

import "OnlineJudge/app/common"

var LavelValidate common.Validator

func init()  {
	rules := map[string]string{
		"lno"			: "required",
		"base"			: "required",
		"cost"			: "required",
		"overdue_cost"  : "required",
		"deposit"		: "required",
	}

	scenes := map[string][]string {
		"add": {"base", "cost", "overdue_cost", "deposit"},
	}
	LavelValidate.Rules = rules
	LavelValidate.Scenes = scenes
}
