package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/qiangxue/fasthttp-routing"
	"golang.org/x/crypto/bcrypt"
)

func (r *UsersRouter) ProcessGetUser(c *routing.Context) (interface{}, error) {
	var sessionUserID int
	session := c.Get("session")
	if session != nil || (session != nil && session.(ISession).ID > 1) {
		sessionUserID = session.(ISession).ID
	}

	queryParams := c.QueryArgs()
	if queryParams.Has("username") && queryParams.Has("id") {
		return nil, errors.New(QueryMultiSearchNotAllowed)
	}

	var data RequestUserData
	err := json.Unmarshal(c.Request.Body(), &data)
	if err != nil {
		return nil, errors.New(BodyError)
	}

	var id int
	if queryParams.Has("id") {
		idParsed, err := queryParams.GetUint("id")
		if err != nil || (err == nil && idParsed < 1) {
			return nil, errors.New(QueryInvalidID)
		}

		id = idParsed
	}

	var username string
	if queryParams.Has("username") {
		username = string(queryParams.Peek("username"))
		if len(username) < 1 || len(username) > 32 {
			return nil, errors.New(QueryUsernameLength)
		}
	}

	fieldsToFetch := []string{}
	for _, v := range data.Fields {
		if v == "" {
			return nil, errors.New(JSONEmptyValuesError)
		}

		if len(v) < 1 || len(v) > 10 {
			return nil, errors.New(QueryInvalidFields)
		}

		allowed := false
		for _, v2 := range AllowedFields {
			if v == v2 {
				allowed = true
				break
			}
		}

		if !allowed {
			return nil, errors.New(QueryInvalidFields)
		}

		if v == "email" && sessionUserID != id {
			return nil, errors.New(UnauthorizedGetUserParams)
		}

		fieldsToFetch = append(fieldsToFetch, v)
	}

	fieldsToString := strings.Join(fieldsToFetch, ", ")
	result, err := r.core.db.Query(fmt.Sprintf("SELECT %s FROM users WHERE id = %d OR username = '%s' LIMIT 1", fieldsToString, id, username))
	if err != nil {
		return nil, errors.New(SQLQueryError)
	}

	columns, err := result.Columns()
	if err != nil {
		return nil, errors.New(SQLQueryError)
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for result.Next() {
		if err := result.Scan(valuePtrs...); err != nil {
			return nil, errors.New(SQLQueryError)
		}
	}

	user := Map{}
	for i, col := range columns {
		switch col {
		case "id":
			user["id"] = values[i].(int64)
		case "username":
			user["username"] = values[i].(string)
		case "email":
			user["email"] = values[i].(string)
		case "gender":
			user["gender"] = values[i].(int64)
		case "created_at":
			user["created_at"] = values[i].(time.Time)
		}
	}

	if len(user) == 0 {
		return nil, errors.New(QueryRecordNotFound)
	}

	return Map{
		"results": user,
	}, nil
}

func (r *UsersRouter) ProcessCreateUser(c *routing.Context) (interface{}, error) {
	var data CreateUserRequest
	err := json.Unmarshal(c.Request.Body(), &data)
	if err != nil {
		return nil, errors.New(BodyError)
	}

	if data.Username == "" || data.Email == "" || data.Password == "" || data.Gender == "" {

		return nil, errors.New(JSONEmptyValuesError)
	}

	if len(data.Password) < 6 || len(data.Password) > 128 {
		return nil, errors.New(PasswordLengthError)
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
		return nil, errors.New(InvalidGenderError)
	}

	result, err := r.core.db.Query(fmt.Sprintf("SELECT FROM users WHERE username = '%s' OR email = '%s' LIMIT 1", data.Username, data.Email))
	if err != nil {
		return nil, errors.New(SQLQueryError)
	}

	if result.Next() {
		return nil, errors.New(UsernameEmailUsedError)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return nil, errors.New(TextToHashError)
	}
	
	// check execution error and rows affected
	exec_result, err := r.core.db.Exec(fmt.Sprintf("INSERT INTO users (username, email, password, gender) VALUES ('%s', '%s', '%s', %d)", data.Username, data.Email, hashedPassword, gender))
	if err != nil {
		return nil, errors.New(SQLExecutionError)
	}

	rowsAffected, _ := exec_result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New(SQLNoRowsAffectedError)
	}

	return RecordedUser, nil
}