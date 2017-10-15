package api

import (
	fmt "fmt"

	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func New() GoGraphQLServer {
	return &server{}
}

func (svc *server) SayHello(ctx context.Context, req *SayHelloReq) (*SayHelloRes, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name cannot be blank")
	}

	msg := fmt.Sprintf("Hello %s", req.Name)

	return &SayHelloRes{
		Message: msg,
	}, nil
}
