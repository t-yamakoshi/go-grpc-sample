package usecase

import (
	"context"
	"fmt"

	"github.com/t-yamakoshi/go-grpc-sample/pkg/domain/repository"
)

var _ IFTodoUsecase = (*TodoUsecase)(nil)

type IFTodoUsecase interface {
	CreateTodo(ctx context.Context, title string, description string) (string, error)
	GetTodo(ctx context.Context, id int64) (*Todo, error)
	UpdateTodo(ctx context.Context, id int64, title string, description string) error
	DeleteTodo(ctx context.Context, id int64) error
}

type TodoUsecase struct {
	TodoRepository repository.IFTodoRepository
}

type Todo struct {
	Id          int64
	Title       string
	Description string
	IsCompleted bool
}

func NewTodoUsecase(todoRepository *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		TodoRepository: todoRepository,
	}
}

func (u *TodoUsecase) CreateTodo(ctx context.Context, title string, description string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (u *TodoUsecase) GetTodo(ctx context.Context, id int64) (*Todo, error) {
	data, err := u.TodoRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Todo{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		IsCompleted: data.IsCompleted,
	}, nil
}

func (u *TodoUsecase) UpdateTodo(ctx context.Context, id int64, title string, description string) error {
	return fmt.Errorf("not implemented")
}

func (u *TodoUsecase) DeleteTodo(ctx context.Context, id int64) error {
	return fmt.Errorf("not implemented")
}
