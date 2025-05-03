package repository

import (
	"context"
	"fmt"

	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/mysqlgen"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/domain/entity"
)

var _ IFTodoRepository = (*TodoRepository)(nil)

type IFTodoRepository interface {
	Create(ctx context.Context, title string, description string) (string, error)
	Get(ctx context.Context, id int64) (*entity.Todo, error)
	Update(ctx context.Context, id int64, title string, description string) error
	Delete(ctx context.Context, id int64) error
}

type Todo struct{}

type TodoRepository struct {
	query *mysqlgen.Query
}

func NewTodoRepository(
	query *mysqlgen.Query,
) *TodoRepository {
	return &TodoRepository{
		query: query,
	}
}

func (r *TodoRepository) Create(ctx context.Context, title string, description string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (r *TodoRepository) Get(ctx context.Context, id int64) (*entity.Todo, error) {
	data, err := r.query.Todo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &entity.Todo{
		Id:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		IsCompleted: data.IsCompleted,
	}, nil
}

func (r *TodoRepository) Update(ctx context.Context, id int64, title string, description string) error {
	return fmt.Errorf("not implemented")
}

func (r *TodoRepository) Delete(ctx context.Context, id int64) error {
	return fmt.Errorf("not implemented")
}
