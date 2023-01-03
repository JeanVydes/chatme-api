package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/qiangxue/fasthttp-routing"
	"golang.org/x/crypto/bcrypt"
)

func (r *AuthRouter) ProcessGetSession(c *routing.Context) (interface{}, error) {
	session := c.Get("session")
	if session == nil || (session != nil && session.(ISession).ID < 1) {
		return nil, errors.New(Unauthorized)
	}

	return session, nil
}


func (r *AuthRouter) ProcessCreateSession(c *routing.Context) (interface{}, error) {
	var data CreateSessionRequest
	err := json.Unmarshal(c.Request.Body(), &data)
	if err != nil {
		return nil, errors.New(BodyError)
	}

	if data.Email == "" || data.Password == "" {
		return nil, errors.New(JSONEmptyValuesError)
	}

	if len(data.Password) < 6 || len(data.Password) > 128 {
		return nil, errors.New(PasswordLengthError)
	}

	rows, err := r.core.db.conn.Query(fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1", data.Email))
	if err != nil {
		return nil, errors.New(SQLQueryError)
	}

	var user IUser
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.Gender, &user.CreatedAt, &user.Username)
		if err != nil {
			return nil, errors.New(SQLQueryError)
		}
	}

	if user.ID < 1 {
		return nil, errors.New(UserNotFoundError)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password))
	if err != nil {
		return nil, errors.New(InvalidPasswordError)
	}

	// Access Token

	session := ISession{
		ID: user.ID,
		CreatedAt: time.Now().Unix(),
		ExpirationDate: time.Now().Add(time.Hour * 24).Unix(),
	}

	sessionString := ParseSession(session)
	aToken, err := Encrypt(hash_key, sessionString)
	if err != nil {
		return nil, errors.New(SessionError)
	}

	return Map{
		"access_token": aToken,
	}, nil
}