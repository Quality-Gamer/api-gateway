package database

import (
	"context"
	"github.com/cheekybits/genny/generic"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func ConnectRedis() *redis.Client {
	return client
}

func SetKey(key string, value generic.Type) {
	_ = client.Set(ctx,key, value, 0).Err()
}

func GetKey(key string) string {
	val, err := client.Get(ctx,key).Result()
	if err != nil {
		return ""
	}
	return val
}

func IncrValue(key string) {
	_ = client.Incr(ctx,key)
}

func HSetKey(key,field string,value generic.Type){
	_ = client.HSetNX(ctx,key,field,value)
}

func HGetKey(key,field string) string {
	return client.HGet(ctx,key,field).Val()
}

func HValKey(key string) string {
	return client.HVals(ctx,key).String()
}

func HDelField(key,field string) {
	_ = client.HDel(ctx,key,field)
}

func HasKey(key string) bool {
	exists, _ := client.Exists(ctx,key).Uint64()

	if exists == 1 {
		return true
	}

	return false
}
