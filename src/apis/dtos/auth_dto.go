package dtos

type LogginRegisterDTO struct {
	Email    string `json:"email" binding:"required,email,lowercase"`
	Password string `json:"password" binding:"required"`
}
