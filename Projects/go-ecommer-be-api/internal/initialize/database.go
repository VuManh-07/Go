package initialize

import (
	"fmt"
	"time"

	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/global"
	"github.com/VuManh-07/Go/Projects/go-ecommer-be-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	// Initialize database connection
	config := global.Configs.Mysql
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	checkErr(err, "Failed to connect to database")

	global.Logger.Info("Database connected successfully")
	global.Mdb = db

	SetPool()

	MigrateDatabase()
}

func SetPool() {
	m := global.Configs.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		global.Logger.Fatal("Failed to get database instance", zap.Error(err))
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(m.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnectionsMaxLifetime))
}

func MigrateDatabase() {
	// Migrate database schema
	err := global.Mdb.AutoMigrate(
		// &YourModel{},
		&po.User{},
		&po.Role{},
	)
	checkErr(err, "Failed to migrate database schema")
	global.Logger.Info("Database migrated successfully")
}

func checkErr(err error, message string) {
	if err != nil {
		global.Logger.Fatal(message, zap.Error(err))
		panic(err)
	}
}
