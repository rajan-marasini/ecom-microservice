package catalog

import (
	"fmt"
	"net"

	"github.com/rajan-marasini/ecom-microservice/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{
		service: s,
	})
	reflection.Register(serv)
	return serv.Serve(list)
}
