package services

import (
	"arayeshyab/src/apis/dto"
	"arayeshyab/src/apis/helpers"
	"arayeshyab/src/configs"
	"arayeshyab/src/databases/schemas"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type tokenService struct{}

func GetTokenService() *tokenService {
	return &tokenService{}
}

func (ts *tokenService) GenerateNewTokens(user *schemas.User, ctx *gin.Context) *helpers.Result {
	// Create a variable to store the tokens we will send back
	var replyToken dto.TokenDTO

	// Generate a new access token for the user
	AccessToken, err := generateAccessToken(user)
	// If there is an error while generating the access token, print the error and return a failure result
	if err != nil {
		fmt.Println(err)                                                                              // Print the error to the console
		return &helpers.Result{Ok: false, Status: 500, Message: "خطا در ساخت توکن دسنرسی", Data: nil} // Return an error message in Persian
	}

	// Generate a new refresh token for the user
	RefreshToken, err := generateRefreshToken(user)
	// If there is an error while generating the refresh token, print the error and return a failure result
	if err != nil {
		fmt.Println(err)                                                                                 // Print the error to the console
		return &helpers.Result{Ok: false, Status: 500, Message: "خطا در ساخت توکن تازه سازی", Data: nil} // Return an error message in Persian
	}

	// Store the generated tokens in our replyToken variable
	replyToken.RefreshToken = RefreshToken
	replyToken.AccessToken = AccessToken

	// Put tokens in client cookies
	ctx.SetCookie("AccessToken", AccessToken, 3600*24*30, "/", "http://localhost:4000", true, true)
	ctx.SetCookie("RefreshToken", RefreshToken, 3600*24*30, "/", "http://localhost:4000", true, true)

	// Return a success result with the tokens
	return &helpers.Result{Ok: true, Status: 201, Message: "خوش امدید", Data: replyToken} // Return a welcome message in Persian
}

func (ts *tokenService) VerifyToken(recived_token string) (*jwt.Token, error) {
	// Get the configuration settings (like secret keys)
	cfg := configs.GetConfigs()

	// Parse the received token to check if it's valid
	token, err := jwt.Parse(recived_token, func(t *jwt.Token) (interface{}, error) {
		// Check if the token is using the correct signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("توکن وارد شده معتبر نیست") // Return an error message in Persian if the token is invalid
		}
		// Return the secret key used to sign the token
		return []byte(cfg.Jwt.AccessTokenKey), nil
	})

	// If there was an error while parsing the token, print it and return an error
	if err != nil {
		fmt.Println(err)                             // Print the error to the console
		return nil, fmt.Errorf("خطا در پردازش توکن") // Return an error message in Persian
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("توکن وارد شده معتبر نیست") // Return an error message in Persian if the token is invalid
	}

	// If everything is good, return the parsed token
	return token, nil
}

func (ts *tokenService) GetTokenClaims(token string) (*dto.TokenDataDTO, *helpers.Result) {
	// Create a new variable to store claims from the token
	claim_result := new(dto.TokenDataDTO)

	// Verify the token to see if it's valid and get its contents
	parsedToken, err := ts.VerifyToken(token)
	// If there was an error verifying the token, return an error result
	if err != nil {
		return nil, &helpers.Result{Ok: false, Status: 400, Message: err.Error(), Data: nil} // Return the error message
	}

	// Try to convert the claims from the parsed token into a specific format
	claims, isConvertionOk := parsedToken.Claims.(jwt.MapClaims)
	// If conversion fails, return an error result saying it's not a valid token
	if !isConvertionOk {
		return nil, &helpers.Result{Ok: false, Status: 400, Message: "رشته وارد شده یک توکن معتبر نیست", Data: nil} // Return an error message in Persian
	}

	// Check if we can get the user ID from the claims
	if ID, isConvertionOk := claims["ID"].(float64); isConvertionOk {
		claim_result.ID = uint(ID) // Store the user ID in our claim_result variable
	} else {
		// If we can't get the ID, return an error result
		return nil, &helpers.Result{Ok: false, Status: 500, Message: "شکست در استخراج اطلاعات از توکن", Data: nil} // Return an error message in Persian
	}
	// If everything is successful, return the claims with a success message
	return claim_result, &helpers.Result{Ok: true, Status: 200, Message: "اطلاعات توکن با موفقیت استخراج شد", Data: nil} // Return success message in Persian
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
