package migrate

import (
	Config "github.com/golang-crud/config"
	Models "github.com/golang-crud/model"
)

//MigrateTable migrando tabelas
func MigrateTable() {
	Config.DB.AutoMigrate(&Models.User{})
}
