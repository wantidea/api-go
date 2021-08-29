package orm

import (
	"api-go/lib/config"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

var sqlDB *sql.DB

func Setup() {
	var err error

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlConfig.UserName,
		config.MysqlConfig.Password,
		config.MysqlConfig.Host,
		config.MysqlConfig.Port,
		config.MysqlConfig.Database)

	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		//SkipDefaultTransaction: true, // 禁用默认事务
	})
	if err != nil {
		log.Fatalf("GORM 数据库连接失败: %v", err)
	}

	sqlDB, err = db.DB()
	if err != nil {
		log.Fatalf("SQL 数据库连接失败: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func CloseDB() {
	_ = sqlDB.Close()
}

func DB() *gorm.DB {
	return db
}

func SqlDB() *sql.DB {
	return sqlDB
}
