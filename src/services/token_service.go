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

func (ts *tokenService) VerifyToken(recived_token string) (*jwt.Token, error) {
	cfg := configs.GetConfigs()

	token, err := jwt.Parse(recived_token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("توکن وارد شده معتبر نیست")
		}
		return []byte(cfg.Jwt.AccessTokenKey), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("خطا در پردازش توکن")
	}

	if !token.Valid {
		return nil, fmt.Errorf("توکن وارد شده معتبر نیست")
	}

	return token, nil
}

func (ts *tokenService) GetTokenClaims(token string) (*dto.TokenDataDTO, *helpers.Result) {
	claim_result := new(dto.TokenDataDTO)

	parsedToken, err := ts.VerifyToken(token)
	if err != nil {
		return nil, &helpers.Result{Ok: false, Status: 400, Message: err.Error(), Data: nil}
	}

	claims, isConvertionOk := parsedToken.Claims.(jwt.MapClaims)
	if !isConvertionOk {
		return nil, &helpers.Result{Ok: false, Status: 400, Message: "رشته وارد شده یک توکن معتبر نیست", Data: nil}
	}

	if id, isConvertionOk := claims["id"].(float64); isConvertionOk {
		claim_result.ID = uint(id)
	} else {
		return nil, &helpers.Result{Ok: false, Status: 500, Message: "شکست در استخراج اطلاعات از توکن", Data: nil}
	}

	return claim_result, &helpers.Result{Ok: true, Status: 200, Message: "اطلاعات توکن با موفقیت استخراج شد", Data: nil}
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
