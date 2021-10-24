package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	fmt.Println("test redis begin:")
	Connect()

	viper.SetDefault("REDIS_REQUEST_NUM", 1)
	req_num := viper.GetInt("REDIS_REQUEST_NUM")

	for i := 0; i < req_num; i++ {
		ExampleClient(i)
		go ExampleClient(i)
		go ExampleClient(i)
		go ExampleClient(i)
		go ExampleClient(i)
	}

}

func Connect() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.SetDefault("REDIS_HOST", "localhost:6379")
	host := viper.GetString("REDIS_HOST")
	fmt.Println("host", host)
	rdb = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func ExampleClient(id int) {

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(id, "key", val, time.Since(start))

	start = time.Now()
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println(id, "key2 does not exist", time.Since(start))
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(id, "key2", val2, time.Since(start))
	}

}
