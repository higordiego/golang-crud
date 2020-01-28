package main

import (
	"fmt"
	"log"

	Config "github.com/golang-crud/config"
	Migrate "github.com/golang-crud/migrate"
	"github.com/golang-crud/router"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	log.Println(Config.BuildDBConfig())
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()

	//setup routes
	r := router.SetupRouter()

	// migrate
	Migrate.MigrateTable()
	// running
	r.Run()
}
