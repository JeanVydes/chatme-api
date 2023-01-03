package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/qiangxue/fasthttp-routing"
)

func ManageAuthorization(c *routing.Context) (bool) {
	authorization := string(c.Request.Header.Peek("Authorization"))
	if len(authorization) == 0 {
		return false
	}

	session, err := Decrypt(hash_key, string(authorization))
	if err != nil {
		return false
	}

	unparsedSession, err := UnParseSession(session)
	if err != nil {
		return false
	}

	if unparsedSession.ExpirationDate < time.Now().Unix() {
		return false
	}

	c.Set("session", unparsedSession)
	c.Set("user-id", unparsedSession.ID)

	return true
}

func AuthMiddleware(next routing.Handler, required bool) routing.Handler {
	return func(c *routing.Context) error {
		authorized := ManageAuthorization(c)
		if !authorized && required {
			Generic(c, 401, Unauthorized, nil, true)
			return nil
		}

		next(c)

		return nil
	}
}

// Access

func UnParseSession(s string) (ISession, error) {
	fragments := strings.Split(s, ".")
	if len(fragments) != 3 {
		return ISession{}, fmt.Errorf("invalid session")
	}

	id, err := strconv.Atoi(fragments[0])
	if err != nil {
		return ISession{}, fmt.Errorf("invalid session")
	}

	createdAt, err := strconv.Atoi(fragments[1])
	if err != nil {
		return ISession{}, fmt.Errorf("invalid session")
	}

	expiration, err := strconv.Atoi(fragments[2])
	if err != nil {
		return ISession{}, fmt.Errorf("invalid session")
	}

	return ISession{
		ID:             id,
		CreatedAt:      int64(createdAt),
		ExpirationDate: int64(expiration),
	}, nil
}

func ParseSession(s ISession) string {
	return fmt.Sprintf("%d.%d.%d", s.ID, s.CreatedAt, s.ExpirationDate)
}