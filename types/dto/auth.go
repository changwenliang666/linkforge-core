package dto

type UserRegistryLoginParams struct {
	Username string `json:"username" binding:"required,max=50"`
	Password string `json:"password" binding:"required,max=255"`
}
