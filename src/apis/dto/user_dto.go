package dto

import "github.com/go-playground/validator/v10"

type UpdateUserDTO struct {
	FullName string `json:"full_name" binding:"required, max=255"`
	UserName string `json:"user_name" binding:"required, max=255, alpha, lowercase"`
	Email    string `json:"email" binding:"required, max=255, email"`
	Address  string `json:"address" binding:"required"`
	Password string `json:"password" binding:"required. min=8"`
}

// if we have validation failed for update user, we'll generate slice and send it to client
func UpdateUserDTO_validation(err error) []string {
	if err.Error() == "EOF" {
		return nil
	}

	var errMsg []string
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Field() {
		case "FullName":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "نام و نام خانوادگی الزامی است")
			} else if err.Tag() == "max" {
				errMsg = append(errMsg, "حداکثر طول نام و نام خانوادگی باید ۲۵۵ حرف باشد")
			}
		case "UserName":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "نام کاربری الزامی است")
			} else if err.Tag() == "max" {
				errMsg = append(errMsg, "حداکثر طول نام کاربری باید ۲۵۵ حرف باشد")
			} else if err.Tag() == "lowercase" || err.Tag() == "alpha" {
				errMsg = append(errMsg, "نام کاربری باید با حروف کوچک انگلیسی باشد")
			}
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
		case "Address":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "ادرس الزامی است")
			}
		case "Password":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "رمز عبور الزامی است")
			}
		}
	}

	return errMsg
}
