package annotation

import (
	"reflect"
)

func GetTableName(pojo interface{}) string {
	v := reflect.ValueOf(pojo).Elem()
	t := v.Type()
	var tablename string = ""
	for i := 0; i < t.NumField(); i++ {
		tablename = t.Field(i).Tag.Get("table")
		if tablename != "" {
			break
		}
	}
	return tablename
}

func ParseTable(pojo interface{}) *Table {
	v := reflect.ValueOf(pojo).Elem()
	t := v.Type()

	var (
		tablename, pk string
	)
	columns := []string{}

	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)

		if _pk := ft.Tag.Get("pk"); _pk != "" {
			pk = _pk
			tablename = ft.Tag.Get("table")
		}

		columns = append(columns, ft.Tag.Get("column"))
	}

	return &Table{Name: tablename, Pk: pk, Columns: columns}
}
