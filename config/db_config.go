package config

import (
	"fmt"
	"time"

	"assemble/logger"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func DBInit(dbConfig DbConfig) {
	DB = LoadDb(dbConfig)
}

func LoadDb(dbConfig DbConfig) *gorm.DB {
	log := logger.DbNew()
	log.SlowThreshold = time.Duration(dbConfig.SlowThreshold) * time.Millisecond
	log.LogLevel = gormlogger.Info
	log.SetAsDefault()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			dbConfig.Host, dbConfig.UserName, dbConfig.DbName, dbConfig.Password, dbConfig.Port),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger:                                   log,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true},
	)

	if err != nil {
		panic("db open fail:" + err.Error())
	}

	sqlDb, err := db.DB()
	if err != nil {
		panic("db init fail:" + err.Error())
	}

	// 连接池最大空闲连接数 (open_counts/2)
	sqlDb.SetMaxIdleConns(dbConfig.IdleCount)
	// 连接池最多同时打开的连接数 (服务器cpu核心数 * 2 + 服务器有效磁盘数)
	sqlDb.SetMaxOpenConns(dbConfig.OpenCount)
	// 连接池连接最大存活时长
	sqlDb.SetConnMaxLifetime(time.Duration(dbConfig.LifeTime) * time.Second)
	// 连接池最大空闲时长
	sqlDb.SetConnMaxIdleTime(time.Duration(dbConfig.IdleTime) * time.Minute)
	return db
}
