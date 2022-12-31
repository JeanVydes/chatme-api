package main

type CreateUserRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Gender string `json:"gender"`
}