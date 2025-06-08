package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"time"
	"weather-api/internal/data"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

const NilErr = redis.Nil

func GetWithoutContext(key string) (*data.Response, bool) {
	cmdBytes, err := Client.Get(context.Background(), key).Bytes()
	data := &data.Response{}
	if err == redis.Nil {
		log.Println("Cache miss")
		return data, false
	} else if err != nil {
		log.Printf("Redis error: %v", err)
	}
	
	err = gob.NewDecoder(bytes.NewReader(cmdBytes)).Decode(data)
	if err != nil {
		log.Printf("Gob error getting: %v", err)
		return data, false
	}
	log.Println("Cache hit")
	return data, true
}

func SetWithoutContext(key string, value *data.Response, expiration time.Duration) {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(value); err != nil {
		log.Printf("Gob error setting %v", err)
	}

	Client.Set(context.Background(), key, b.Bytes(), expiration)
}

func RedisInit() {
	Client = redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // No password set
        DB:		  0,  // Use default DB
        Protocol: 2,  // Connection protocol
    })

	if _, err := Client.Ping(context.Background()).Result(); err != nil {
		log.Fatal("Failed to load Redis")
	} else {
		log.Println("Redis initialized successfully")
	}
}