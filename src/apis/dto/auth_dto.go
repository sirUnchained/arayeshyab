package dto

import "github.com/go-playground/validator/v10"

type AuthDTO struct {
	Email    string `json:"email" binding:"required,max=255,email,lowercase"`
	Password string `json:"password" binding:"required,min=8"`
}

// if we have validation failed for login/register, we'll generate slice and send it to client
func AuthDTO_GenerateFailedMap(err error) []string {
	if err.Error() == "EOF" {
		return nil
	}

	var errMsg []string
	for _, err := range err.(validator.ValidationErrors) {

		switch err.Field() {
		case "Email":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "ایمیل الزامی است")
			} else if err.Tag() == "email" {
				errMsg = append(errMsg, "ایمیل معتبر نیست")
			} else if err.Tag() == "lowercase" {
				errMsg = append(errMsg, "ایمیل باید با حروف انگلسی کوچک باشد")
			} else if err.Tag() == "max" {
				errMsg = append(errMsg, "حداکثر طول ایمیل باید ۲۵۵ حرف باشد")
			}
		case "Password":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "رمز عبور الزامی است")
			} else if err.Tag() == "min" {
				errMsg = append(errMsg, "رمز عبور باید حداقل ۸ حرف باشد")
			}
		}

	}

	return errMsg
}
