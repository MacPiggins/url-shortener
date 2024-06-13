package redis

import (
	"context"
	"url-shortener/internal/database"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

func (db *RedisDB) Set(token, link string) error {
	exists, err := db.client.Exists(context.Background(), token).Result()
	if err != nil {
		return err
	}

	if exists == 1 {
		return database.UniqueError{}
	}

	err = db.client.Set(context.Background(), token, link, 0).Err()	
	return err
}

func (db *RedisDB) Get(token string) (string, error) {
	link, err := db.client.Get(context.Background(), token).Result()
	if err != nil {
		if err == redis.Nil {
			return "", database.NotFoundError{}
		}
		return "", err
	}
	return link, nil
}

func (db *RedisDB) Close() {
	db.client.Close()
}

func New(link string) (*RedisDB, error) {
	db := RedisDB{}
	var err error

	opts, err := redis.ParseURL(link)
	if err != nil {
		return nil, err
	}
	db.client = redis.NewClient(opts)
	return &db, err
}
