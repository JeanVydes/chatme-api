package main

import (
	"fmt"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"github.com/qiangxue/fasthttp-routing"
)

type UsersRouter struct {
	core *Core
	router *routing.RouteGroup
} 

func (r *UsersRouter) SetRoutes() {
	r.SetEndpoints()
}

func (r *UsersRouter) SetEndpoints() {
	r.router.Get("", r.GetUser) // Get user
	r.router.Post("", r.CreateUser) // Create new user record in SQL database
}

func (r *UsersRouter) GetUser(c *routing.Context) error {
	c.Write([]byte("<user>"))

	return nil
}

func (r *UsersRouter) CreateUser(c *routing.Context) error {
	var data CreateUserRequest
	err := json.Unmarshal(c.Request.Body(), &data)
	if err != nil {
		Generic(c, 400, BodyError, nil, true)
		return nil
	}

	if data.Username == "" || data.Email == "" || data.Password == "" || data.Gender == "" {
		Generic(c, 400, JSONEmptyValuesError, nil, true)
		return nil
	}

	if len(data.Password) < 6 || len(data.Password) > 128 {
		Generic(c, 400, PasswordLengthError, nil, true)
		return nil
	}

	var gender int8
	switch data.Gender {
	case "female":
		gender = Female
	case "male":
		gender = Male
	case "other":
		gender = Other
	default:
		Generic(c, 400, InvalidGenderError, nil, true)
		return nil
	}

	result, err := r.core.db.Query(fmt.Sprintf("SELECT FROM users WHERE username = '%s' OR email = '%s' LIMIT 1", data.Username, data.Email))
	if err != nil {
		Generic(c, 500, SQLQueryError, nil, true)
		return nil
	}

	if result.Next() {
		Generic(c, 400, UsernameEmailUsedError, nil, true)
		return nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		Generic(c, 500, TextToHashError, nil, true)
		return nil
	}
	
	// check execution error and rows affected
	exec_result, err := r.core.db.Exec(fmt.Sprintf("INSERT INTO users (username, email, password, gender) VALUES ('%s', '%s', '%s', %d)", data.Username, data.Email, hashedPassword, gender))
	if err != nil {
		Generic(c, 500, SQLExecutionError, nil, true)
		return nil
	}

	rowsAffected, _ := exec_result.RowsAffected()
	if rowsAffected == 0 {
		Generic(c, 500, SQLNoRowsAffectedError, nil, true)
		return nil
	}

	Generic(c, 200, RecordedUser, nil, true)
	return nil
}