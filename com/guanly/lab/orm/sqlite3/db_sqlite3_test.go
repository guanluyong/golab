package sqlite3

import (
	"testing"
)

func TestGetDb(t *testing.T) {
	conf := map[string]string{
		"driver_name":      "sqlite3",
		"data_source_name": "orm.db",
		"max_open_conns":   "5",
		"max_idle_conns":   "3",
	}
	db := GetDb(conf)

	if err := db.Ping(); err != nil {
		t.Error("Sqlite3 GetDb test failed.")
	}
}
