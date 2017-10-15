package api

import (
	fmt "fmt"

	context "golang.org/x/net/context"
)

type server struct{}

func New() GoGraphQLServer {
	return &server{}
}

func (svc *server) SayHello(ctx context.Context, req *SayHelloReq) (*SayHelloRes, error) {
	msg := fmt.Sprintf("Hello %s", req.Name)

	return &SayHelloRes{
		Message: msg,
	}, nil
}
