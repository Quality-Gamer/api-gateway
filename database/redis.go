package database

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"github.com/go-redis/redis"
	"os"
)

var client = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_HOST"),
	Password: os.Getenv("REDIS_PASS"), // no password set
	DB:       0,  // use default DB
})

func ConnectRedis() *redis.Client {
	return client
}

func SetKey(key string, value generic.Type) {
	_ = client.Set(key, value, 0).Err()
}

func GetKey(key string) string {
	fmt.Println("DEBUG REDIS")
	fmt.Println(client.Get("gw:ms:ky:qg"))
	val, err := client.Get(key).Result()
	if err != nil {
		return ""
	}
	return val
}

func IncrValue(key string) {
	_ = client.Incr(key)
}

func HSetKey(key,field string,value generic.Type){
	_ = client.HSetNX(key,field,value)
}

func HGetKey(key,field string) string {
	return client.HGet(key,field).Val()
}

func HValKey(key string) string {
	return client.HVals(key).String()
}

func HDelField(key,field string) {
	_ = client.HDel(key,field)
}

func HasKey(key string) bool {
	exists := client.Exists(key)

	if exists.Val() == 1 {
		return true
	}

	return false
}
