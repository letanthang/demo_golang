package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func main() {
	fmt.Println("test redis begin:")
	ExampleClient()
}

func ExampleClient() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("REDIS_HOST", "localhost:6379")
	host := viper.GetString("REDIS_HOST")
	fmt.Println("host", host)
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("key", val, time.Since(start))

	start = time.Now()
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist", time.Since(start))
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2, time.Since(start))
	}

}
