package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/wojtekw92/go-playground/pkg/protocol/grpc"
	"github.com/wojtekw92/go-playground/pkg/service/v1"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string

}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}


	v1API := v1.NewGreeterServer()

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}