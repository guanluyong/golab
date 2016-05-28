package sqlite3

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

func GetDb(conf map[string]string) *sql.DB {
	driverName, ok := conf["driver_name"]
	if !ok {
		log.Fatal("Not found driver_name attr.")
		panic("Not found driver_name attr.")
	}

	dataSourceName, ok := conf["data_source_name"]
	if !ok {
		log.Fatal("Not found data_source_name attr.")
		panic("Not found data_source_name attr.")
	}

	xdb, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("Open db error.")
		panic(err)
	}

	if moc, ok := conf["max_open_conns"]; ok {
		if _moc, err := strconv.Atoi(moc); err == nil {
			xdb.SetMaxOpenConns(_moc)
		} else {
			log.Fatal(err.Error())
		}
	}
	if mic, ok := conf["max_idle_conns"]; ok {
		if _mic, err := strconv.Atoi(mic); err == nil {
			xdb.SetMaxIdleConns(_mic)
		} else {
			log.Fatal(err.Error())
		}
	}

	/*if err := xdb.Ping(); err != nil {
		panic(err)
	}*/

	return xdb
}
