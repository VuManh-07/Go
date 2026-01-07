package initialize

import (
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/global"
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Configs.Logger)
}
