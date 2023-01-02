package main

type CreateSessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Gender string `json:"gender"`
}

type GetSessionRequest struct {
	Token string `json:"token"`
}