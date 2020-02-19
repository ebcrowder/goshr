package db

import (
	"fmt"
	"log"

	"github.com/ebcrowder/goshr/schema"
	"github.com/go-redis/redis/v7"
)

type Redis struct {
	DB *redis.Client
}

func ConnectRedis() (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(pong)

	return &Redis{client}, nil
}

func (r *Redis) Insert(file *schema.File) (string, error) {

	err := r.DB.Set(file.Name, file.Key, 0).Err()
	if err != nil {
		panic(err)
	}

	return "hi", nil
}

func (r *Redis) Delete(id int) error {
	return nil
}

func (r *Redis) GetFiles() ([]schema.File, error) {

	var fileList []schema.File

	val, err := r.DB.Get("test right here").Result()
	if err != nil {
		panic(err)
	}

	fileList[0] = [Name "test"]schema.File

	return fileList, nil
}
