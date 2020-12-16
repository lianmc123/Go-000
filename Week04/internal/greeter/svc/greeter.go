package svc

import (
	"context"
	pb "directory/api/greeter"
	"directory/internal/greeter/biz"
)

type GreeterSvc struct {
	biz *biz.Biz
}

func (s *GreeterSvc) Hello(_ context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	response, err := s.biz.Hello(request.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloResponse{
		Msg: response,
	}, nil
}

func (s *GreeterSvc) HelloAgain(_ context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	response, err := s.biz.HelloAgain(request.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloResponse{
		Msg: response,
	}, nil
}

func NewGreeterSvc(biz *biz.Biz) pb.GreeterServer {
	return &GreeterSvc{
		biz: biz,
	}
}
