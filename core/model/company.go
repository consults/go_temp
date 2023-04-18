package model

import (
	log "github.com/sirupsen/logrus"
	"sync_pid_redis/core/db"
)

var TableCompany = "company"
var CompanyDao = companyDao{}

type companyDao struct {
}

type Company struct {
	Cid string `json:"cid"`
	Pid string `json:"pid"`
}
type CompanyName struct {
	Name string `json:"name"`
}

// 获取全部pid存在但是current_cid 不存在的数据,
func (c *companyDao) GetAllPid() *[]*Company {
	log.Info("查询公司current_id 为空，但是pid不为空的数据")
	var companys []*Company
	mysql := db.GetMysql()
	mysql.Table(TableCompany).Where("pid IS NOT NULL AND current_cid IS NULL").Find(&companys)
	return &companys
}

func (c companyDao) GetAllName() *[]*CompanyName {
	log.Infoln("获取全部pid不存在，current_id 也不存在的数据，只获取name")
	var companyNames []*CompanyName
	mysql := db.GetMysql()
	mysql.Table(TableCompany).Where("WHERE pid IS NULL AND current_cid IS NOT NULL").Find(&companyNames)
	return &companyNames
}
