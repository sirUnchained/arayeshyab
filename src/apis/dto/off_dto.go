package dto

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

// todo we have some problem with validate time i have to fix it
type CreateOffDto struct {
	Amount    uint      `json:"amount" binding:"required,numeric,min=0,max=100"`
	Code      string    `json:"code" binding:"required,len=16"`
	ExpiresAt time.Time `json:"expires_at" binding:"required"`
}

func CreateOffDto_validate(err error) []string {
	if err.Error() == "EOF" {
		return nil
	}

	fmt.Printf("%+v\n", err)

	errMsg := []string{}

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Field() {
		case "Amount":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "مقدار تخفیف الزامی است")
			} else if err.Tag() == "numeric" {
				errMsg = append(errMsg, "مقدار تخفیف باید عددی و به درصد باشد")
			} else if err.Tag() == "min" || err.Tag() == "max" {
				errMsg = append(errMsg, "تخفیف حداکثر ۱۰۰ و حداقل ۰ باید باشد")
			}
		case "Code":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "کد تخفیف الزامی است")
			} else if err.Tag() == "len" {
				errMsg = append(errMsg, "طول کد تخفیف باید ۱۶ حرف باشد")
			}
		case "ExpiresAt":
			if err.Tag() == "required" {
				errMsg = append(errMsg, "زمان انقضای کد تخفیف الزامی است")
			}
		}

	}

	return errMsg
}
