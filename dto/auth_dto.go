package dto

type UserRegistry struct {
	Username string `json:"username" binding:"required,max=20"`
	Password string `json:"password" binding:"required,max=255"`
}
