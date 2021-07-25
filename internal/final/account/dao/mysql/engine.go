package mysql

import (
	"go-demo/internal/final/account/dao"
	"go-demo/internal/final/account/define"
	"go-demo/pkg/mysql"

	"github.com/pkg/errors"
)

type MyEngine struct {
	*mysql.MyEngine
}

func NewAngInitMyEngine(opts ...mysql.Opt) (dao.AcountEngine, error) {
	e := &MyEngine{
		MyEngine: mysql.NewEngine(opts...),
	}
	err := e.Init()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return e, nil
}

func (m *MyEngine) DeleteAcount(id string) error {
	sql := `delete from comment where id=?`

	return m.Exec(sql, id)
}

func (m *MyEngine) AddAcount(data *define.AcountData) error {
	sql := `insert into comment values(?,?,?,?)`

	return m.Exec(sql, data.AcountID, data.Name, data.Age, data.EmailAddr)
}

func (m *MyEngine) QueryAcount(acountID string, limit, offset int) ([]*define.AcountData, error) {
	sql := `select * from comment where acount_id=? limit ?,?`

	db := m.GetDb()
	stmt, err := db.Prepare(sql)
	defer stmt.Close()

	rows, err := stmt.Query(acountID, limit, offset)
	if err != nil {
		return nil, errors.Wrapf(err, "QueryAcount:%s acountID:%s limit:%S offset:%s", sql, acountID, limit, offset)
	}
	defer rows.Close()

	datas := make([]*define.AcountData, 0, offset)
	for rows.Next() {
		data := new(define.AcountData)

		err := rows.Scan(&data.AcountID, &data.Name, &data.Age, &data.EmailAddr)
		if err != nil {
			return nil, errors.Wrapf(err, "QueryAcount:%s acountID:%s limit:%S offset:%s", sql, acountID, limit, offset)
		}

		datas = append(datas, data)
	}

	return datas, nil
}
