package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mvc/pkg/config"
)

func NewMysqlConn(cfg *config.Database) *gorm.DB {
	// set database dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Charset,
		cfg.ParseTime,
	)

	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 禁用默认复数表名
			SingularTable: true,
		},
	}

	// open connection
	mysqlDb, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		panic("database open error! " + err.Error())
	}
	db, _ := mysqlDb.DB()

	db.SetMaxIdleConns(cfg.MaxIdleCons)
	db.SetMaxOpenConns(cfg.MaxOpenCons)

	return mysqlDb
}
