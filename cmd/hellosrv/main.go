package main

import (
	"log"
	"net"

	"github.com/iheanyi/go-grpc-graphql-simple-example/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error listening: %v", err)
	}

	srv := grpc.NewServer()
	helloSvc := api.New()
	api.RegisterGoGraphQLServer(srv, helloSvc)
	log.Print("Starting up the server")
	log.Printf("starting up server: %v", srv.Serve(lis))
}
