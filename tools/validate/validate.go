package validate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

type Validator struct {
	Rules  map[string]string
	Scenes map[string][]string
}

func (validator *Validator) ValidateMap(m map[string]interface{}, scene string) error {
	if _, ok := validator.Scenes[scene]; !ok {
		return errors.New("scene is not exists")
	}
	v := validate.Map(m)

	//  添加规则
	for _, field := range validator.Scenes[scene] {
		v.StringRule(field, validator.Rules[field])
	}

	if v.Validate() {
		return nil
	}
	return errors.New(v.Errors.One())
}

func (validator *Validator) ValidateRequest(c *gin.Context, scene string) error {
	// 判断scene是否存在
	if _, ok := validator.Scenes[scene]; !ok {
		return errors.New("scene is not exists")
	}
	httpData, err := validate.FromRequest(c.Request)
	if err != nil {
		return err
	}
	// 创建验证器
	v := httpData.Create()
	//  添加规则
	for _, field := range validator.Scenes[scene] {
		v.StringRule(field, validator.Rules[field])
	}

	if v.Validate() {
		return nil
	}
	return errors.New(v.Errors.One())

}
