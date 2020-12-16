//+build wireinject

package di

import (
	"directory/internal/greeter/biz"
	"directory/internal/greeter/dao"
	"directory/internal/greeter/svc"
	"directory/pkg/database/sql"
	"github.com/google/wire"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.NewDao, biz.NewBiz, sql.NewConn, newApp, svc.NewGreeterSvc))
}
