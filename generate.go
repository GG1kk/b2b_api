// +build generate

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go

//go:generate protoc -I. --go_out=paths=source_relative,plugins=grpc:. ./rpc/catalog-service.proto
//go:generate protoc -I. --go_out=paths=source_relative,plugins=grpc:. ./rpc/setting-service.proto
//go:generate protoc -I. --go_out=paths=source_relative,plugins=grpc:. ./rpc/users-service.proto
//go:generate protoc -I. --go_out=paths=source_relative,plugins=grpc:. ./rpc/wallet-service.proto
//go:generate protoc -I. --go_out=paths=source_relative,plugins=grpc:. ./rpc/trading-service.proto

package pkg
