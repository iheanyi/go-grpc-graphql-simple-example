package api

//go:generate protoc -I=.:$GOPATH/src/github.com/gogo/protobuf/protobuf --proto_path=${GOPATH}/src:. --gogofaster_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types:. api.proto
