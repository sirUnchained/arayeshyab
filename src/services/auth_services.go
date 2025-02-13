package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"

	"github.com/gin-gonic/gin"
)

type authServices struct{}

func GetAuthServices() *authServices {
	return &authServices{}
}

func (ah *authServices) Login(ctx *gin.Context) *helpers.Result {
	userData := new(dto.AuthDTO)
	if err := ctx.ShouldBindBodyWithJSON(userData); err != nil {
		errs := dto.AuthDTO_GenerateFailedMap(err)

		return &helpers.Result{Ok: false, Status: 400, Message: "اعتبار سنجی شکست خورد لطفا ورودی هارا با دقت وارد کنید", Data: errs}
	}

	return &helpers.Result{Ok: true, Status: 200, Message: "خوش آمدید", Data: nil}
}

func (ah *authServices) Register(ctx *gin.Context) *helpers.Result {
	userData := new(dto.AuthDTO)
	if err := ctx.ShouldBindBodyWithJSON(userData); err != nil {
		errs := dto.AuthDTO_GenerateFailedMap(err)

		return &helpers.Result{Ok: false, Status: 400, Message: "اعتبار سنجی شکست خورد لطفا ورودی هارا با دقت وارد کنید", Data: errs}
	}

	return &helpers.Result{Ok: true, Status: 201, Message: "خوش آمدید", Data: nil}
}
