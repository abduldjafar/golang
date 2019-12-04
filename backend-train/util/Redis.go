package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func newClient(db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       db, // use default DB
	})

	return client
}

func SaveToRedisMovies(email, data string, db int) {
	client := newClient(db)
	err := client.Set(email, data, time.Hour*24).Err()
	if err != nil {
		panic(err)
	}
}

func GetDataFromRedis(email string, data int, db int) string {
	client := newClient(db)
	value, err := client.Get(email).Result()
	if err == redis.Nil {
		fmt.Println(nil)
	} else {
		panic(err)
	}
	return value
}
