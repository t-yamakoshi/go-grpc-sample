package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/grpcgen"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/config"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/service"
)

func main() {
	config := config.NewConfig()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Printf("Error starting server: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	todoService := service.NewTodoService()

	grpcgen.RegisterTodoServiceServer(s, todoService)

	go func() {
		log.Printf("Starting server on port %d", config.Port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()
}
