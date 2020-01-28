package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// DB intancia de gorm.DB
var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig ponteiro de configuração.
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		DBName:   "golang",
		Password: "root",
	}
	return &dbConfig
}

// DbURL função de string para url do banco de dados
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
