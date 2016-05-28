package dbconfig

import (
	"testing"
)

func TestGetConf(t *testing.T) {
	m := GetConf("db.config")
	if m["driver_name"] != "sqlite3" {
		t.Error("driver_name is not correct!")
	}
}
