package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/configs"
	"arayeshyab/src/databases/schemas"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct{}

func GetTokenService() *tokenService {
	return &tokenService{}
}

func (ts *tokenService) GenerateNewTokens(user *schemas.User) *helpers.Result {
	var replyToken dto.TokenDTO

	AccessToken, err := generateAccessToken(user)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطا در ساخت توکن دسنرسی", Data: nil}
	}

	RefreshToken, err := generateRefreshToken(user)
	if err != nil {
		fmt.Println(err)
		return &helpers.Result{Ok: false, Status: 500, Message: "خطا در ساخت توکن تازه سازی", Data: nil}
	}

	replyToken.RefreshToken = RefreshToken
	replyToken.AccessToken = AccessToken
	return &helpers.Result{Ok: true, Status: 201, Message: "خوش امدید", Data: replyToken}
}

func generateAccessToken(user *schemas.User) (string, error) {
	cfg := configs.GetConfigs()
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID": user.ID,
		// "exp": time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpirePerMinute)).Unix(),
	})

	tokenStr, err := tokenObj.SignedString([]byte(cfg.Jwt.AccessTokenKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// todo => at the end of project i'll add refresh token feature
func generateRefreshToken(user *schemas.User) (string, error) {
	cfg := configs.GetConfigs()
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID": user.ID,
		// "exp": time.Now().Add(time.Duration(cfg.Jwt.AccessTokenExpirePerMinute)).Unix(),
	})

	tokenStr, err := tokenObj.SignedString([]byte(cfg.Jwt.AccessTokenKey))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
