package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

/*
Код делает очень простую вещь.
Функция init() автоматически вызывается Go, код извлекает информацию о соединении из .env файла,
затем строит строку соединения и использует её для соединения с базой данных.
*/
func init() {

	e := godotenv.Load() // Загрузка файла .env
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s  sslmode=disable", dbHost, username, dbName, password) // Создать строку подключения
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})

}

func GetDB() *gorm.DB {
	return db
}
