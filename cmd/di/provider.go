package di

import (
	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/gateway/mysql"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/config"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/domain/repository"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/usecase"
)

type Repository struct {
	TodoRepository *repository.TodoRepository
}

type Usecase struct {
	TodoUsecase *usecase.TodoUsecase
}

type Provider struct {
	Config     *config.Config
	Usecase    Usecase
	Repository Repository
}

func NewProvider() (*Provider, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	mysqlConfig := mysql.NewConfig(cfg)
	db, err := mysql.NewMySQLDB(mysqlConfig)
	if err != nil {
		return nil, err
	}
	gormClient := mysql.NewGormClient(db)
	todoRepository := repository.NewTodoRepository(
		gormClient,
	)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	usecase := Usecase{
		TodoUsecase: todoUsecase,
	}
	repository := Repository{
		TodoRepository: todoRepository,
	}
	return &Provider{
		Usecase:    usecase,
		Repository: repository,
	}, nil
}
