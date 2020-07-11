package auth

import (
	"fmt"
	"github.com/Quality-Gamer/api-gateway/conf"
	"github.com/Quality-Gamer/api-gateway/database"
)

func KeyValidator(key string) bool {
	fmt.Println(conf.GetAuthKey())
	token := database.GetKey(conf.GetAuthKey())
	fmt.Println(key)
	fmt.Println(token)
	return key == token
}
