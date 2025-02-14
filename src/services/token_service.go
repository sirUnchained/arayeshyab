package services

type tokenService struct{}

func GetTokenService() *tokenService {
	return &tokenService{}
}
