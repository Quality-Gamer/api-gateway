package auth

import (
	"conf"
	"database"
)

func KeyValidator(key string) bool {
	token := database.GetKey(conf.GetAuthKey())
	return key == token
}
