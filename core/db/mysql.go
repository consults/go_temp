package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync_pid_redis/core/config"
)

var mySqlDb *gorm.DB

func GetMysql() *gorm.DB {
	return mySqlDb
}
func InitMysql() {
	// 初始化mysql数据库
	log.Info("开始初始化mysql数据库")
	conf := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.MysqlUser,
		conf.MysqlPassword, conf.MysqlIp, conf.MysqlPort, conf.MysqlDb)
	log.Infoln(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("连接数据库失败")
		panic(err.Error())
	}
	mySqlDb = db
	log.Info("连接数据库成功")
}
