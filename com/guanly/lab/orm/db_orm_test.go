package orm

import (
	//an "guanly/lab/orm/annotation"
	"testing"
)

func TestNewDefaultOrmPanel(t *testing.T) {
	op := NewDefaultOrmPanel(1)
	if !op.Ping() {
		t.Error("Can't ping database")
	}
}

func TestNewOrmPanel(t *testing.T) {
	op := NewOrmPanel("dbconfig/db.config", 1)
	if !op.Ping() {
		t.Error("Can't ping database")
	}
}

type User struct {
	Id   int64  `column:"id" pk:"id" table:"user"`
	Name string `column:"name" length:16`
}

func TestGetTable(t *testing.T) {
	op := NewDefaultOrmPanel(1)

	pojo := &User{Id: 1, Name: "Jack"}
	tb, err := op.GetTable(pojo)
	if err != nil {
		t.Error("Get Table failed.")
	}
	t.Log(tb)
}
