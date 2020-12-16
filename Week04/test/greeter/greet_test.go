package greeter

import (
	"context"
	pb "directory/api/greeter"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
	"testing"
)

func TestGreetSvc(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	response, err := client.HelloAgain(context.Background(), &pb.HelloRequest{
		Name: "傻子A",
	})
	if err != nil {
		fromError, ok := status.FromError(err)
		if ok {
			t.Error(fromError.Code())
			t.Error(fromError.Message())
		} else {
			t.Error(err)
		}
		return
	}
	fmt.Println(response.Msg)

	response, err = client.Hello(context.Background(), &pb.HelloRequest{
		Name: "傻子B",
	})
	if err != nil {
		fromError, ok := status.FromError(err)
		if ok {
			t.Error(fromError.Code())
			t.Error(fromError.Message())
		} else {
			t.Error(err)
		}
		return
	}
	fmt.Println(response.Msg)
}
