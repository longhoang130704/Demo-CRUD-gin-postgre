package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=@Hlong04 dbname=golang_database port=5432 sslmode=disable"
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Không thể kết nối đến cơ sở dữ liệu:", err)
    }

    DB = database
    fmt.Println("Kết nối PostgresSQL thành công!")
}