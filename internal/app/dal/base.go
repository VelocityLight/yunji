package dal

import (
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"yunji/configs"
)

func (conn MysqlInfo) RawWrapper(sql string, values ...interface{}) (tx *gorm.DB) {
	if strings.Contains(sql, "@") || strings.Contains(sql, "?") {
		return conn.DB.Raw(sql, values...)
	} else {
		return conn.DB.Raw(sql)
	}
}

// Mysql handler infomation
type MysqlInfo struct {
	DB *gorm.DB
	// Anything else...
}

var DBConn = &MysqlInfo{}

func Connect(config *configs.ConfigYaml) {
	// Connect
	conn, err := gorm.Open(mysql.Open(config.Secret.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 600)

	DBConn.DB = conn

	// conn.AutoMigrate()
	// Close(Delayed)
	// defer db.Close()
}
