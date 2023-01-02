package main

import "time"

const (
	Female = iota
	Male
	Other
)

const (
	AccessToken = "Access"
	AccessExpirationTime = time.Hour * 24
)

const (
	Created = "created"
	Authorized = "authorized"
	Unauthorized = "unauthorized"
	RecordedUser 		   = "recorded_user"
)

const (
	BodyError              = "empty_body_or_invalid_json"
	JSONEmptyValuesError   = "empty_values"
	PasswordLengthError    = "password_length"
	InvalidGenderError     = "invalid_gender"
	UsernameEmailUsedError = "username_email_used"
	UserNotFoundError 	   = "user_not_found"
	InvalidPasswordError   = "invalid_password"
	SessionError		  = "session_error"
	InvalidTokenError 	= "invalid_token"
	InternalServerError = "internal_server_error"
	// Technical errors
	TextToHashError        = "hash_0001"
	SQLExecutionError      = "db_0001"
	SQLQueryError          = "db_0002"
	SQLNoRowsAffectedError = "db_0005"
)

type IHTTPMessage struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"exited_code"`
}

type IUser struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	Gender    int8   `json:"gender"`
	CreatedAt string `json:"created_at"`
}

type ISession struct {
	ID             int   `json:"user_id"`
	CreatedAt      int64 `json:"created_at"`
	ExpirationDate int64 `json:"expiration"`
}