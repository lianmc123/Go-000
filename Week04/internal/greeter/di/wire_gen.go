// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"directory/internal/greeter/biz"
	"directory/internal/greeter/dao"
	"directory/internal/greeter/svc"
	"directory/pkg/database/sql"
)

// Injectors from wire.go:

func InitApp() (*App, func(), error) {
	db, cleanup, err := sql.NewConn()
	if err != nil {
		return nil, nil, err
	}
	helloRepo := dao.NewDao()
	bizBiz := biz.NewBiz(db, helloRepo)
	greeterServer := svc.NewGreeterSvc(bizBiz)
	app, cleanup2, err := newApp(greeterServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return app, func() {
		cleanup2()
		cleanup()
	}, nil
}
