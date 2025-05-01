package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	schema "github.com/t-yamakoshi/go-grpc-sample/schema/mysql"
)

func main() {
	schemas := []any{
		&schema.Todo{},
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

	director := mysql.Open("root:@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(director)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	g.UseDB(db)
	db.AutoMigrate(schemas...)

	g.ApplyBasic(schemas...)
	g.ApplyInterface(func() {}, schemas...)

	g.Execute()
}
