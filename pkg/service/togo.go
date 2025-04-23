package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/grpcgen"
)

var _ grpcgen.TodoServiceServer = (*TodoService)(nil)

type TodoService struct {
	grpcgen.UnimplementedTodoServiceServer
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) CreateTodo(context.Context, *grpcgen.CreateTodoRequest) (*grpcgen.CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (s *TodoService) GetTodo(context.Context, *grpcgen.GetTodoRequest) (*grpcgen.GetTodoResponse, error) {
	return &grpcgen.GetTodoResponse{
		Todo: &grpcgen.Todo{
			Id:          "1",
			Title:       "test",
			Description: "test",
			IsCompleted: false,
		},
	}, nil
}
func (s *TodoService) UpdateTodo(context.Context, *grpcgen.UpdateTodoRequest) (*grpcgen.UpdateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodo not implemented")
}
func (s *TodoService) DeleteTodo(context.Context, *grpcgen.DeleteTodoRequest) (*grpcgen.DeleteTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodo not implemented")
}
func (s *TodoService) ListTodos(context.Context, *grpcgen.ListTodosRequest) (*grpcgen.ListTodosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTodos not implemented")
}

func (s *TodoService) mustEmbedUnimplementedTodoServiceServer() {}
