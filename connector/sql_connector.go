package connector

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type SqlConnector struct {
	*gorm.DB
}

type SqlConnectorOption struct {
	DBName   string
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	SSLMode  string
}

func NewSqlConnector(option SqlConnectorOption) (*gorm.DB, error) {
	args := fmt.Sprintf(`host=%s port=%s user=%s dbname=%s password=%s sslmode=%s`, option.Host, option.Port, option.User, option.DBName, option.Password, option.SSLMode)
	return gorm.Open("postgres", args)
}
