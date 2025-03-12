package dto

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token *string `json:"token"`
}

type RegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	EMail    string `json:"email"`
}

type RegisterResponse struct {
	Token *string `json:"token"`
}
