package infrastructure

import (
	"fmt"
	"go-proj/go-proj/utils"

	// go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SetupModels : initializing mysql database
func SetupModels() *gorm.DB {
	get := utils.GetEnvWithKey
	USER := get("DB_USER")
	PASS := get("DB_PASS")
	HOST := get("DB_HOST")
	PORT := get("DB_PORT")
	DBNAME := get("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}

	return db
}
