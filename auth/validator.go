package auth

import (
	"gateway/conf"
	"gateway/database"
)

func KeyValidator(key string) bool {
	token := database.GetKey(conf.GetAuthKey())
	return key == token
}
