package validate

import "OnlineJudge/app/common"

var StaffValidate common.Validator

func init()  {
	rules := map[string]string{
		"sno"			: "required",
		"name"			: "required",
		"sex"			: "required",
		"number"		: "required",
		"licence"		: "required|len:18",
	}

	scenes := map[string][]string {
		"add": {"sex", "name", "number", "licence"},
		"findByName": {"name"},
		"findByID"	: {"customer_id"},
		"findByLicence"	: {"licence"},
	}
	StaffValidate.Rules = rules
	StaffValidate.Scenes = scenes
}