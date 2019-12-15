package data

import (
	"fmt"
	"gwi-challenge/common"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

// InitDB initializes our databse and returns its referance
func InitDB() *gorm.DB {

	if common.IsDemo() {
		err := os.Remove("./gwi_demo.db")
		if err != nil {
			fmt.Println("db deletion error: ", err)
		}
	}

	db, err := gorm.Open("sqlite3", common.GetConfig().DbPath)
	if err != nil {
		fmt.Println("db creation error: ", err)
	}

	db.DB().SetMaxIdleConns(10)
	database = db

	createTables()

	return database
}

// GetDB returns the reference of our database
func GetDB() *gorm.DB {
	return database
}

func createTables() {
	if common.IsDemo() {
		fileContent, err := ioutil.ReadFile(common.GetConfig().SqlDropPath)
		if err != nil {
			fmt.Print(err)
			return
		}
		database.Exec(string(fileContent))
	}
	fileContent, err := ioutil.ReadFile(common.GetConfig().SqlCreatePath)
	if err != nil {
		fmt.Print(err)
		return
	}
	database.Exec(string(fileContent))
}

func PopulateDbWithDemoData() {
	fileContent, err := ioutil.ReadFile(common.GetConfig().SqlPopulatePath)
	if err != nil {
		fmt.Print(err)
		return
	}
	database.Exec(string(fileContent))
}
