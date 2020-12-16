## 目录结构

```text
├── README.md
├── api
│     └── greeter
│         ├── greeter.pb.go
│         └── greeter.proto     #pb文件
├── cmd
│     └── greeter-service
│         └── main.go           #程序主入口
├── configs                     #配置文件
│     ├── db.yaml
│     └── grpc.yaml
├── go.mod
├── go.sum
├── internal
│     └── greeter
│         ├── biz
│         │     └── biz.go
│         ├── dao
│         │     └── dao.go
│         ├── di
│         │     ├── app.go
│         │     ├── wire.go     #依赖反转生成服务入口
│         │     └── wire_gen.go
│         ├── model
│         │     └── models.go  
│         └── svc
│             └── greeter.go    # svc层
├── pkg
│     └── database
│         └── sql
│             └── mysql.go      # 数据库连接
└── test
    └── greeter
        └── greet_test.go       # 测试
```