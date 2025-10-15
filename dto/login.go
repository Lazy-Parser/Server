package dto

type LoginFirstDTO struct {
	Username string `json:"username" binding:"required"`
}

type LoginSecondDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
