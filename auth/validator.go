package auth

func KeyValidator(key string) bool {
	token := "2cfd648bde02d8c3271ae1d9f7226f7e" //temporário
	return key == token
}
