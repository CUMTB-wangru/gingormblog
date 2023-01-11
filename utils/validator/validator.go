package validator

// 验证规则包--包含各种语言提示信息
import (
	"fmt"
	"ginblog-master/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// 全局验证
func Validate(data interface{}) (string, int) {
	// 实例化
	validate := validator.New()
	// 翻译包实例化
	uni := unTrans.New(zh_Hans_CN.New())
	// 引入指定翻译器
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	// 注册翻译方法
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	// 配置好想要输出的映射  label
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	// 不知道传入的数据类型时，需要使用类型断言
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
