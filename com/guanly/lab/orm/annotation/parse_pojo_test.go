package annotation

import (
	"testing"
)

type User struct {
	Id   int64  `pk:"id" table:"user" column:"id"`
	Name string `column:"name" length:"16"`
}

func TestGetTableName(t *testing.T) {
	u := &User{Name: "Jack"}
	tablename := GetTableName(u)

	if tablename == "" {
		t.Error("Not found tablename")
	}
}

func TestParseTable(t *testing.T) {
	u := &User{Id: 1, Name: "Steven"}
	tb := ParseTable(u)
	if tb == nil {
		t.Error("Can't parse pojo")
	}

}
