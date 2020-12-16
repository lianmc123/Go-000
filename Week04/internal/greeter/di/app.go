package di

import (
	"context"
	pb "directory/api/greeter"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type App struct {

}

func newApp(svc pb.GreeterServer) (app *App, closeFunc func(), err error) {
	grpcViper := viper.New()
	grpcViper.AddConfigPath(".")
	grpcViper.SetConfigFile("grpc.yaml")
	if err = grpcViper.ReadInConfig(); err != nil {
		return
	}
	listen, err := net.Listen("tcp", grpcViper.GetString("server.addr"))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, svc)
	go func() {
		if err := server.Serve(listen); err != nil {
			panic(err)
		}
	}()
	closeFunc = func() {
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		server.Stop()
		cancel()
	}
	return
}
