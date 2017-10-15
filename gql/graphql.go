package gql

import (
	"context"

	"github.com/iheanyi/go-graphql-experiment/api"
)

var Schema = `
	schema {
		query: Query
		# mutation: Mutation
	}
	type Query {
		hello(name: String!): String!
	}
	# type Mutation {}
`

type Resolver struct {
	helloSvc api.GoGraphQLClient
}

func NewResolver(helloClient api.GoGraphQLClient) *Resolver {
	return &Resolver{
		helloSvc: helloClient,
	}
}

func (r *Resolver) Hello(args struct{ Name string }) (string, error) {
	request := &api.SayHelloReq{
		Name: args.Name,
	}

	res, err := r.helloSvc.SayHello(context.TODO(), request)
	if err != nil {
		return "", err
	}

	return res.Message, nil
}
