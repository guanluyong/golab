package dbconfig

import (
	"guanly/lab/proputil"
)

const (
	CONF_PATH string = "dbconfig/db.config"
)

func GetConf(confPath string) map[string]string {
	if confPath == "" {
		confPath = CONF_PATH
	}
	conf := proputil.GetProps(confPath)
	return conf
}
