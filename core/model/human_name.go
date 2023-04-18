package model

import (
	"sync_pid_redis/core/db"
)

var TableHumanName = "human_name"

type HumanName struct {
	Hid       string `json:"hid"`
	PersionId string `json:"persion_id"`
}

// 获取全部humanname的数据
func GetAllHumanName() *[]HumanName {
	var humannames []HumanName
	mysql := db.GetMysql()
	mysql.Table(TableHumanName).Where("persion_id != '' AND persion_id IS NOT NULL AND persion_id != 'Null'").Find(&humannames)
	return &humannames
}

func GetAllHumanNameCount() int64 {
	var count int64
	mysql := db.GetMysql()
	mysql.Table(TableHumanName).Where("persion_id != '' AND persion_id IS NOT NULL AND persion_id != 'Null'").Count(&count)
	return count
}
