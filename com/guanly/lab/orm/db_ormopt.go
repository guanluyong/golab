package orm

import (
	"errors"
	an "guanly/lab/orm/annotation"
)

func (op *OrmPanel) GetTable(pojo interface{}) (*an.Table, error) {
	tableName := an.GetTableName(pojo)
	if tableName == "" {
		return nil, errors.New("Not found Table name at Pojo")
	}

	if t, ok := op.Tables[tableName]; ok {
		return t, nil
	}

	table := an.ParseTable(pojo)
	op.Tables[tableName] = table

	return table, nil
}

func (op *OrmPanel) GetParams(pojo interface{}) []interface{} {
	params := make([]interfaces{})

	// parse pojo value
	return params
}

func (op *OrmPanel) save(pojo interface{}) (int64, error) {
	tn, err := op.GetTable(pojo)
	if err != nil {
		return nil, err
	}

	table := op.Tables[tn]
	sql := table.InsertSQL()

	stmt, err := op.Db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(op.GetParams(pojo))
	if err != nil {
		return nil, err
	}

	id := res.LastInsertId()
	return id, nil
}

func (op *OrmPanel) Ping() bool {
	if err := op.Db.Ping(); err != nil {
		return false
	}
	return true
}
