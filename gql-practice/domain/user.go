package domain

import (
	"math/rand"
	"time"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"username"`
}

const (
	USER_ID_TKNR1216   = "tknr1216"
	USER_ID_TAKASHI003 = "takashi003"
	USER_ID_TARO115    = "taro115"
	USER_ID_HANAKO43   = "hanako43"
	USER_ID_YAMADA234  = "yamada234"
)

func GenerateUserId() string {
	userIds := []string{
		USER_ID_TKNR1216,
		USER_ID_TAKASHI003,
		USER_ID_TARO115,
		USER_ID_HANAKO43,
		USER_ID_YAMADA234,
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userIds[r.Intn(len(userIds))]

}
