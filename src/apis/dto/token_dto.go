package dto

type TokenDTO struct {
	AccessToken  string `binding:"required"`
	RefreshToken string `binding:"required"`
}
