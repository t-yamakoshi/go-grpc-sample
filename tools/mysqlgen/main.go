package main

import (
	"log"

	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/t-yamakoshi/go-grpc-sample/schema/mysql"
)

type Querier interface {
	// SELECT * FROM @@table WHERE id = @id
	GetByID(id int64) (gen.T, error)
}

func main() {
	schemas := []any{
		&mysql.Todo{},
	}
	g := gen.NewGenerator(
		gen.Config{
			OutPath:           "pkg/adapter/mysqlgen",
			Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			FieldNullable:     true,
		},
	)

	director := mysqlDriver.Open("root:@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(director)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	g.UseDB(db)
	if err := db.AutoMigrate(schemas...); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	g.ApplyBasic(schemas...)
	g.ApplyInterface(func(Querier) {}, schemas...)

	g.Execute()
}
