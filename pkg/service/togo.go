package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/grpcgen"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/usecase"
)

var _ grpcgen.TodoServiceServer = (*TodoService)(nil)

type TodoService struct {
	grpcgen.UnimplementedTodoServiceServer
	usecase usecase.IFTodoUsecase
}

func NewTodoService(
	todoUsecase *usecase.TodoUsecase,
) *TodoService {
	return &TodoService{
		usecase: todoUsecase,
	}
}

func (s *TodoService) CreateTodo(context.Context, *grpcgen.CreateTodoRequest) (*grpcgen.CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (s *TodoService) GetTodo(ctx context.Context, req *grpcgen.GetTodoRequest) (*grpcgen.GetTodoResponse, error) {
	res, err := s.usecase.GetTodo(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get todo: %v", err)
	}
	return &grpcgen.GetTodoResponse{
		Todo: &grpcgen.Todo{
			Id:          res.Id,
			Title:       res.Title,
			Description: res.Description,
			IsCompleted: res.IsCompleted,
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
