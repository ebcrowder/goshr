package db

import (
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

	return &Redis{client}, nil
}

func (r *Redis) Insert(file *schema.File) (string, error) {
	err := r.DB.HMSet(file.ID, []string{"name", file.Name, "key", file.Key}).Err()
	if err != nil {
		panic(err)
	}

	return file.ID, nil
}

func (r *Redis) Delete(id string) error {
	err := r.DB.Del(id).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *Redis) GetFiles(id string) ([]interface{}, error) {
	val, err := r.DB.HMGet(id, "name", "key").Result()
	if err != nil {
		panic(err)
	}

	return val, nil
}
