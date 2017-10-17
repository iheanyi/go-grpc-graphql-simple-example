package main

import (
	"log"
	"net/http"
	"time"

	"github.com/iheanyi/go-graphql-experiment/api"
	"github.com/iheanyi/go-graphql-experiment/gql"
	graphql "github.com/neelance/graphql-go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error dialing grpc server: %v", err)
	}
	defer conn.Close()
	helloSvc := api.NewGoGraphQLClient(conn)

	resolver := gql.NewResolver(helloSvc)
	schema := graphql.MustParseSchema(gql.Schema, resolver)
	mux := gql.NewRouter(schema)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
