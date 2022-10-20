package dal

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"yunji/configs"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	// Params
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		config.TiDB.UserName,
		config.TiDB.PassWord,
		config.TiDB.Host,
		config.TiDB.Port,
		config.TiDB.DataBase,
		config.TiDB.CharSet,
		config.TiDB.TimeZone,
	)

	// Connect
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{
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
