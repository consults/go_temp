package model

import (
	"sync_pid_redis/core/db"
)

var TableCategory = "category_code"

type CategoryCodeDao struct {
}
type CategoryCode struct {
	CategoryCode string `json:"category_code"`
	CateThird    string `json:"cate_third"`
	CateSecond   string `json:"cate_second"`
	CateFirst    string `json:"cate_first"`
}

// 获取全部category_code的数据
func (c *CategoryCodeDao) GetAllCategoryData() *[]*CategoryCode {
	var categoryCodes []*CategoryCode
	mysql := db.GetMysql()
	mysql.Table(TableCategory).Find(&categoryCodes)
	return &categoryCodes
}
