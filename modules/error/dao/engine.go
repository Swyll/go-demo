package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type myEngine struct{}

func NewEngine() *myEngine {
	return &myEngine{}
}

func (m *myEngine) GetData(id string) (*data, error) {
	sql := "select * from tablename where id=?"

	sqlOpen()
	prepared(sql)

	d, err := m.query(id)
	if isNoRowErr(err) {
		return new(data), nil
	}
	if err != nil {
		return nil, errors.WithStack(fmt.Errorf(fmt.Sprintf("sql:%s--id:%s", sql, id)+"--err:%w", err))
	}

	return d, nil
}

func sqlOpen()            {}
func prepared(sql string) {}

func (m *myEngine) query(params ...interface{}) (*data, error) {
	return nil, sql.ErrNoRows
}

func isNoRowErr(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
