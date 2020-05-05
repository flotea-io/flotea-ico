package global

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

type Form struct {
	NameSurname string `form:"nameSurname" valid:"Required"`
	Email       string `form:"email" valid:"Email"`
	Phone       string `form:"phone"`
	Message     string `form:"message" valid:"Required"`
	Recaptcha   string `form:"g-recaptcha-response" valid:"Required"`
}

func init() {
	validation.SetDefaultMessage(map[string]string{
		"Required": "Pole nie może być puste",
		"Email":    "Adres e-mail musi być prawidłowy",
	})
}

func Validate(u Form) map[string][]*validation.Error {
	valid := validation.Validation{}
	_, err := valid.Valid(&u)
	if err != nil {
		fmt.Println(err)
	}
	return valid.ErrorsMap
}

func EmptyForm() Form {
	var u Form
	return u
}
