package db

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type MyRedis struct {
	Addr     string
	Password string
	DB       int
	redis    *redis.Client
}

func (r *MyRedis) BatSetHashData(redis_key string, dataList *[]*map[string]interface{}) {
	pipe := r.redis.Pipeline()
	for _, data := range *dataList {
		err := pipe.HMSet(redis_key, *data).Err()
		if err != nil {
			log.Error(data)
			panic(err)
		}
	}
	// 执行管道命令
	_, err := pipe.Exec()
	if err != nil {
		log.Error("批量插入数据失败")
		panic(err)
	}
	log.Infoln("批量插入redis成功")
}
func (r *MyRedis) CloseRedis() {
	err := r.redis.Close()
	if err != nil {
		log.Error("关闭redis失败")
		panic(err.Error())
	}
	log.Warningf("关闭数据库成功")
}
func (r *MyRedis) InitRedis() {
	rd := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	r.redis = rd
	log.Warningf("初始化redis成功")
}
