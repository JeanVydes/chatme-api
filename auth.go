package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/qiangxue/fasthttp-routing"
)

func ManageAuthorization(c *routing.Context) error {
	authorization := string(c.Request.Header.Peek("Authorization"))
	if len(authorization) == 0 {
		Generic(c, 401, Unauthorized, nil, true)
		return nil
	}

	session, err := Decrypt(hash_key, string(authorization))
	if err != nil {
		Generic(c, 401, Unauthorized, nil, true)
		return nil
	}

	unparsedSession, err := UnParseSession(session)
	if err != nil {
		Generic(c, 401, Unauthorized, nil, true)
		return nil
	}

	if unparsedSession.ExpirationDate < time.Now().Unix() {
		Generic(c, 401, Unauthorized, nil, true)
		return nil
	}

	c.Set("session", unparsedSession)

	return nil
}

func AuthMiddleware(next routing.Handler) routing.Handler {
	return func(c *routing.Context) error {
		ManageAuthorization(c)
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