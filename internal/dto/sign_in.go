package dto

type SignInInput struct {
	Password string `json:"password"`
}

type SignInOutput struct {
	Token string `json:"token"`
}
