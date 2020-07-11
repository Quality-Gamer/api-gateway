package auth

import (
	"github.com/Quality-Gamer/api-gateway/conf"
	"github.com/Quality-Gamer/api-gateway/database"
)

func KeyValidator(key string) bool {
	token := database.GetKey(conf.GetAuthKey())
	return key == token
}
