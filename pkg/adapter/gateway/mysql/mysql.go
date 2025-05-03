package mysql

import (
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	mysqlDriver "github.com/go-sql-driver/mysql"

	"github.com/t-yamakoshi/go-grpc-sample/pkg/adapter/mysqlgen"
	"github.com/t-yamakoshi/go-grpc-sample/pkg/config"
)

var (
	params = map[string]string{
		"parseTime":    "true",
		"loc":          url.QueryEscape("UTC"),
		"timeout":      "3s",
		"collation":    "utf8mb4_bin",
		"writeTimeout": "60s",
	}
)

func NewConfig(cfg *config.Config) *mysql.Config {
	mysqlDriver := mysqlDriver.NewConfig()
	mysqlDriver.User = cfg.MYSQL_USER
	mysqlDriver.Passwd = cfg.MYSQL_PASSWORD
	mysqlDriver.Net = "tcp"
	mysqlDriver.Addr = cfg.MYSQL_HOST + ":" + cfg.MYSQL_PORT
	mysqlDriver.DBName = cfg.MYSQL_DB
	mysqlDriver.Params = params
	mysqlDriver.ParseTime = true
	mysqlDriver.Collation = "utf8mb4_bin"
	return &mysql.Config{
		DSNConfig:  mysqlDriver,
		DriverName: "mysql",
		DSN:        mysqlDriver.FormatDSN(),
		Conn:       nil,
	}
}

func NewMySQLDB(cfg *mysql.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(*cfg), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewGormClient(db *gorm.DB) *mysqlgen.Query {
	return mysqlgen.Use(db)
}
