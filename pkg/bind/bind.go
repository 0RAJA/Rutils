package bind

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

/*
业务接口校验:
	required	必填
	gt			大于
	gte			大于等于
	lt			小于
	lte			小于等于
	min			最小值
	max			最大值
	oneof		参数集内的其中之一
	len			长度要求与 len 给定的一致
在结构体中应用到了两个 tag 标签，分别是 form 和 binding，它们分别代表着表单的映射字段名和入参校验的规则内容，其主要功能是实现参数绑定和参数检验。
*/

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() (ret []string) {
	for _, v := range v {
		ret = append(ret, v.Error())
	}
	return
}

func (v *ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

/*
BindAndValid
在上述代码中，我们主要是针对入参校验的方法进行了二次封装，
在 BindAndValid 方法中，通过 ShouldBind 进行参数绑定和入参校验
*/
func BindAndValid(c *gin.Context, v interface{}) (ok bool, errs ValidErrors) {
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans") //翻译
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for k, v := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     k,
				Message: v,
			})
		}
		return false, errs
	}
	return true, nil
}

func FormatBindErr(errs ValidErrors) string {
	return fmt.Sprintf("BindAndValid err: %v", errs)
}
