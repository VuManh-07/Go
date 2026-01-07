package initialize

import "github.com/VuManh-07/Go/Projects/go-ecommer-be-api/global"

func Run() {
	InitConfig()
	InitLogger()
	global.Logger.Info("Logger initialized successfully")
	InitDatabase()
	InitRedis()

	r := InitRouter()

	r.Run(":8081")
}
