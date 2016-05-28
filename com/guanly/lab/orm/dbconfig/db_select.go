package dbconfig

import (
	"database/sql"
	"errors"
	"guanly/lab/orm/sqlite3"
)

func SelectDB(confPath string, kind DbKind) (db *sql.DB, err error) {
	switch int(kind) {
	case 0:
		err = errors.New("Invaild DB")
	case 1:
		db = sqlite3.GetDb(GetConf(confPath))
	case 2:
		err = errors.New("Not suport Mysql db")
	case 3:
		err = errors.New("Not suport Oracle db")
	case 4:
		err = errors.New("Not suport Db2 db")
	default:
		err = errors.New("Invaild DB")
	}
	return
}
