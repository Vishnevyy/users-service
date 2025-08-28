package grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	userpb "github.com/Vishnevyy/project-protos/proto/user"
	"github.com/Vishnevyy/users-service/internal/user"
)

func RunGRPC(svc *user.Service) error {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, NewHandler(svc))

	return s.Serve(l)
}
