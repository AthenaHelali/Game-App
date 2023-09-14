package param

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User  UserInfo `json:"user"`
	Token Tokens   `json:"token"`
}
