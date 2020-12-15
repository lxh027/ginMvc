package validate

import "OnlineJudge/app/common"

var CustomerValidate common.Validator

func init()  {
	rules := map[string]string{
		"customer_id"	: "required",
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
	CustomerValidate.Rules = rules
	CustomerValidate.Scenes = scenes
}

