package main

import (
	"fmt"
	"log"

	schema "github.com/t-yamakoshi/go-grpc-sample/schema/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

//go generate go run main.go

func main() {
	g := gen.NewGenerator(
		gen.Config{
			OutPath:           "../pkg/adapter/mysqlgen",
			Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			FieldNullable:     true,
		},
	)

	config := mysql.Open("mysql:mysqlroot@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(config)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	mysqlDb, err := db.DB()
	if err != nil {
		log.Fatal("failed to get database:", err)
	}
	defer mysqlDb.Close()

	stmt := &gorm.Statement{DB: db}
	if err = stmt.Parse(&schema.Todo{}); err != nil {
		log.Fatal("failed to parse statement:", err)
	}
	fmt.Println(stmt.Schema.CreateTableSQL())

	for _, field := range stmt.Schema.Fields {
		for _, tag := range field.Tag {
			if tag.Name == "gorm" {
				value, ok := tag.Options["uniqueIndex"]
				if ok {
					if ok {
						fmt.Printf("CREATE UNIQUE INDEX idx_%s_%s ON %s(%s);\n", stmt.Schema.Table, field.DBName, stmt.Schema.Table, field.DBName)
					}
				}
			}
		}

		err = db.AutoMigrate(&schema.Todo{})
		if err != nil {
			log.Fatal("failed to migrate database:", err)
		}
	}
}
