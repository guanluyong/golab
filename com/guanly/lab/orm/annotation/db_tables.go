package annotation

import (
	"fmt"
)

type Table struct {
	Name    string
	Pk      string
	Columns []string
}

func (t *Table) String() string {
	return fmt.Sprintf("Table name: %s, Primary key: %s, Columns: %v", t.Name, t.Pk, t.Columns)
}

func (t *Table) InsertSQL() string {
	sql := "INSERT INTO " + t.Name + " VALUES ("
	// for k, v := range t.Columns {

	// }
	return sql
}
