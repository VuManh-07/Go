package global

import (
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/pkg/logger"
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/pkg/setting"
	"github.com/redis/go-redis/v9"

	"gorm.io/gorm"
)

var (
	Configs setting.Config
	Logger  *logger.LoggerZap
	Rdb     *redis.Client
	Mdb     *gorm.DB
)
