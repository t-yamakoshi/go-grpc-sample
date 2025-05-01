package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/t-yamakoshi/go-grpc-sample/cmd/di"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/grpcgen"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/config"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/service"
)

func main() {
	provider, err := di.NewProvider()
	if err != nil {
		log.Printf("Error initializing provider: %v", err)
		os.Exit(1)
	}

	config, err := config.NewConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.APP_PORT))
	if err != nil {
		log.Printf("Error starting server: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()

	todoService := service.NewTodoService(
		provider.Usecase.TodoUsecase,
	)

	grpcgen.RegisterTodoServiceServer(s, todoService)

	go func() {
		log.Printf("Starting server on port %d", config.APP_PORT)
		reflection.Register(s)
		if err := s.Serve(listener); err != nil {
			log.Printf("Error serving: %v", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()
}
