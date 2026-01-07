package initialize

import (
	"fmt"

	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/global"
)

func InitDatabase() {
	// Initialize database connection
	config := global.Configs.Mysql
	fmt.Printf("Connecting to database at %s:%d\n", config.Host, config.Port)
}
