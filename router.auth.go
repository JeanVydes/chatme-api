package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/qiangxue/fasthttp-routing"
	"golang.org/x/crypto/bcrypt"
)

type AuthRouter struct {
	core *Core
	router *routing.RouteGroup
} 

func (r *AuthRouter) SetRoutes() {
	r.SetEndpoints()
}

func (r *AuthRouter) SetEndpoints() {
	r.router.Get("/session", AuthMiddleware(r.GetSession, true))
	r.router.Post("/session", r.CreateSession)
}

func (r *AuthRouter) GetSession(c *routing.Context) error {
	result, err := r.ProcessGetSession(c)
	if err != nil {
		Generic(c, 500, err.Error(), nil, true)
		return nil
	}

	Generic(c, 200, "OK", result, true)

	return nil
}

func (r *AuthRouter) CreateSession(c *routing.Context) error {
	var data CreateSessionRequest
	err := json.Unmarshal(c.Request.Body(), &data)
	if err != nil {
		Generic(c, 400, BodyError, nil, true)
		return nil
	}

	if data.Email == "" || data.Password == "" {
		Generic(c, 400, JSONEmptyValuesError, nil, true)
		return nil
	}

	if len(data.Password) < 6 || len(data.Password) > 128 {
		Generic(c, 400, PasswordLengthError, nil, true)
		return nil
	}

	rows, err := r.core.db.conn.Query(fmt.Sprintf("SELECT * FROM users WHERE email = '%s' LIMIT 1", data.Email))
	if err != nil {
		Generic(c, 500, SQLQueryError, nil, true)
		return nil
	}

	var user IUser
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.Gender, &user.CreatedAt, &user.Username)
		if err != nil {
			Generic(c, 500, SQLQueryError, nil, true)
			return nil
		}
	}

	if user.ID < 1 {
		Generic(c, 404, UserNotFoundError, nil, true)
		return nil
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password))
	if err != nil {
		Generic(c, 400, InvalidPasswordError, nil, true)
		return nil
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
		Generic(c, 500, SessionError, nil, true)
		return nil
	}

	Generic(c, 200, Created, Map{
		"access_token": aToken,
	}, true)

	return nil
}