package dto

import "github.com/go-playground/validator/v10"

type AuthDTO struct {
	Email    string `json:"email" binding:"required,email,lowercase"`
	Password string `json:"password" binding:"required,min=8"`
}

// if we have validation failed for login/register, we'll generate slug and send it to client
func AuthDTO_GenerateFailedMap(err error) []string {
	var errMsg []string

	for _, err := range err.(validator.ValidationErrors) {

		switch err.Field() {
		case "Email":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "ایمیل الزامی است")
			}
			if err.Tag() == "email" {
				errMsg = append(errMsg, "ایمیل معتبر نیست")
			}
			if err.Tag() == "lowercase" {
				errMsg = append(errMsg, "ایمیل باید با حروف انگلسی کوچک باشد")
			}
		case "Password":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "رمز عبور الزامی است")
			}
			if err.Tag() == "min" {
				errMsg = append(errMsg, "رمز عبور باید حداقل ۸ حرف باشد")
			}
		}

	}

	return errMsg
}
