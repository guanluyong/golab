package orm

import (
	"database/sql"
	an "guanly/lab/orm/annotation"
	"guanly/lab/orm/dbconfig"
)

type OrmPanel struct {
	Db     *sql.DB
	Tables map[string]*an.Table
}

func NewDefaultOrmPanel(kind dbconfig.DbKind) *OrmPanel {
	return NewOrmPanel("", kind)
}

func NewOrmPanel(confPath string, kind dbconfig.DbKind) *OrmPanel {
	db, err := dbconfig.SelectDB(confPath, kind)
	if err != nil {
		panic(err)
	}

	return &OrmPanel{Db: db, Tables: make(map[string]*an.Table)}
}
