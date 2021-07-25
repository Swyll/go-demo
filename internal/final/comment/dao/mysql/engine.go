package mysql

import (
	"go-demo/internal/final/comment/dao"
	"go-demo/internal/final/comment/define"
	"go-demo/pkg/mysql"

	"github.com/pkg/errors"
)

type MyEngine struct {
	*mysql.MyEngine
}

func NewAngInitMyEngine(opts ...mysql.Opt) (dao.CommentEngine, error) {
	e := &MyEngine{
		MyEngine: mysql.NewEngine(opts...),
	}
	err := e.Init()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return e, nil
}

func (m *MyEngine) DeleteComment(id string) error {
	sql := `delete from comment where id=?`

	return m.Exec(sql, id)
}

func (m *MyEngine) AddComment(data *define.CommentData) error {
	sql := `insert into comment values(?,?,?)`

	return m.Exec(sql, data.ID, data.AcountID, data.Comment)
}

func (m *MyEngine) QueryComment(acountID string, limit, offset int) ([]*define.CommentData, error) {
	sql := `select * from comment where acount_id=? limit ?,?`

	db := m.GetDb()
	stmt, err := db.Prepare(sql)
	defer stmt.Close()

	rows, err := stmt.Query(acountID, limit, offset)
	if err != nil {
		return nil, errors.Wrapf(err, "QueryComment:%s acountID:%s limit:%S offset:%s", sql, acountID, limit, offset)
	}
	defer rows.Close()

	datas := make([]*define.CommentData, 0, offset)
	for rows.Next() {
		data := new(define.CommentData)

		err := rows.Scan(&data.ID, &data.AcountID, &data.Comment)
		if err != nil {
			return nil, errors.Wrapf(err, "QueryComment:%s acountID:%s limit:%S offset:%s", sql, acountID, limit, offset)
		}

		datas = append(datas, data)
	}

	return datas, nil
}
