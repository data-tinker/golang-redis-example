package main

import (
	"log"

	"encoding/json"
	"golang-redis-example/models"

	"github.com/go-redis/redis"
)

const (
	name = "Alex"
	age  = 31
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, pingErr := client.Ping().Result()
	log.Println(pong, pingErr)

	nameSetErr := client.Set("name", name, 0).Err()
	if nameSetErr != nil {
		log.Fatal(nameSetErr)
	}

	nameVal, nameGetErr := client.Get("name").Result()
	log.Println(nameVal, nameGetErr)

	userJSON, marshalErr := json.Marshal(models.User{
		Name: name,
		Age:  age,
	})
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	userSetErr := client.Set("user", userJSON, 0).Err()
	if userSetErr != nil {
		log.Fatal(userSetErr)
	}

	userVal, userGetErr := client.Get("user").Result()
	log.Println(userVal, userGetErr)
}
