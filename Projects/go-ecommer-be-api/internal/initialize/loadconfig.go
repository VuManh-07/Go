package initialize

import (
	"fmt"

	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/global"
	"github.com/spf13/viper"
)

func InitConfig() {
	// Load configuration settings
	viper.AddConfigPath("./config/") // optionally look for config in the working directory
	viper.SetConfigName("local")     // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// Read values from the config file
	fmt.Println("Server Port:", viper.GetString("server.port"), "MySQL Port:", viper.GetInt("mysql.port"))

	if err := viper.Unmarshal(&global.Configs); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}
