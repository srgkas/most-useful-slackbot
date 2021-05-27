package internal

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/srgkas/most-useful-slackbot/internal/config"
)

var r *redis.Client
var ctx = context.Background()

func InitStorage() {
	if r == nil {
		return
	}

	conf := config.GetConfig().GetRedisConf()

	r = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
		DB: conf.DbNumber,
	})
}

func Set(key, value string) error {
	return r.Set(ctx, key, value, 0).Err()
}

func Get(key, value string) (string, error) {
	return r.Get(ctx, key).Result()
}