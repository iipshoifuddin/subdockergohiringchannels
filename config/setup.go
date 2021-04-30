// models/setup.go
package config

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	URL := fmt.Sprintf(os.Getenv("MYSQL_CONN_STRING"))
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
