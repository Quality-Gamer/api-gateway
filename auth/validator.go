package auth

import (
	"api-gateway/conf"
	"api-gateway/database"
)

func KeyValidator(key string) bool {
	token := database.GetKey(conf.GetAuthKey())
	return key == token
}
