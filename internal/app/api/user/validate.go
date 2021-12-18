package user

import "mvc/tools/validate"

var Validate validate.Validator

func init() {
	rules := map[string]string{
		"id":       "required",
		"username": "required|maxLen:25",
		"password": "required|minLen:6|maxLen:16",
	}

	scenes := map[string][]string{
		"addUser": {"username", "password"},
		"login":   {"username", "password"},
	}
	Validate.Rules = rules
	Validate.Scenes = scenes
}
