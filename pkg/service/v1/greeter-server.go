package v1

import (
	"context"
	"log"

	"github.com/wojtekw92/go-playground/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type greeterServer struct {
}

// NewToDoServiceServer creates ToDo service
func NewGreeterServer() v1.GreeterServer {
	return &greeterServer{}
}

// SayHello implements helloworld.GreeterServer
func (s *greeterServer) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}