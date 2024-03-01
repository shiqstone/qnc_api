package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormopentracing "gorm.io/plugin/opentracing"

	"qnc/biz/mw/viper"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	config := viper.Conf.DB
	DB, err = gorm.Open(mysql.Open(config.MysqlDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}
