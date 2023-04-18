package main

import (
	log "github.com/sirupsen/logrus"
	"sync_pid_redis/core/config"
	"sync_pid_redis/core/logger"
	"time"
)

func main() {
	logger.InitLog("log", "log", 1024*1024*100, time.Hour*24*7)
	config.Init()
	log.Infoln("进行业务开发")

}
