package dbconfig

import (
	"testing"
)

func TestSelectDb1(t *testing.T) {
	db, err := SelectDB("db.config", 1)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(db)
}

func TestSelectDb2(t *testing.T) {
	_, err := SelectDB("db.config", 2)
	if err != nil {
		t.Log(err.Error())
	} else {
		t.Error("Should not support")
	}

}
