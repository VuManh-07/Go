package initialize

import (
	"context"
	"fmt"

	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	config := global.Configs.Redis
	fmt.Println("Redis Config:", config)
	// Initialize Redis connectionx
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis:", zap.Error(err))
		panic(err)
	}

	global.Rdb = rdb
	global.Logger.Info("Redis connected successfully")
}
