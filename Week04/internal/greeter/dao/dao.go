package dao

import (
	"database/sql"
	"directory/internal/greeter/biz"
)

type GreetDao struct{}

func NewDao() biz.HelloRepo {
	return new(GreetDao)
}

func (d *GreetDao) SaveHello(db *sql.DB, hello *biz.HelloItem) error {
	_, err := db.Exec("insert into hello(`message`, `source`) values (?, ?)", hello.Message, hello.Source)
	return err
}
