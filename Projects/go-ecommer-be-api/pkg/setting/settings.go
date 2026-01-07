package setting

type Config struct {
	Mysql  DatabaseConfig `mapstructure:"mysql"`
	Logger LoggerConfig   `mapstructure:"logger"`
}

type DatabaseConfig struct {
	Host                   string `mapstructure:"host"`
	Port                   int    `mapstructure:"port"`
	Username               string `mapstructure:"username"`
	Password               string `mapstructure:"password"`
	DBName                 string `mapstructure:"dataname"`
	MaxIdleConnections     int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections     int    `mapstructure:"maxOpenConnections"`
	ConnectionsMaxLifetime int    `mapstructure:"connectionsMaxLifetime"`
}

type LoggerConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxBackups int    `mapstructure:"maxBackups"`
	MaxAge     int    `mapstructure:"maxAge"`
}
