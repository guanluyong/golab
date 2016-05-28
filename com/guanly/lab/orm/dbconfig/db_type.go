package dbconfig

import (
	"strconv"
)

type DbKind uint

const (
	Invalid DbKind = iota
	Sqlite3
	Mysql
	Oracle
	Db2
)

var DbKindNames = []string{
	"invalid", "sqlite3", "mysql", "oracle", "db2",
}

func (d DbKind) String() string {
	if int(d) < len(DbKindNames) {
		return DbKindNames[d]
	}
	return "DbKind" + strconv.Itoa(int(d))
}
