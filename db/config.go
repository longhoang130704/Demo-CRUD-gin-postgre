package db

import (
	"GIN-TUTORIAL/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    // dsn 1 use to run localhost; dsn 2 use to run in Docker container
    // comment line 19 to run localhost
    // comment line 18 to run in Docker
    //  dsn := "host=localhost user=postgres password=12345678 dbname=golang_database port=5432 sslmode=disable"
    dsn := "host=postgres user=postgres password=12345678 dbname=gin_database port=5432 sslmode=disable"

    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Không thể kết nối đến cơ sở dữ liệu:", err)
    }

    DB = database
    fmt.Println("Kết nối PostgresSQL thành công!")
}

func MigrateDatabase(db *gorm.DB) {
    var user models.Users
    err := db.AutoMigrate(&user)
    if err != nil {
        log.Fatalf("Không thể tự động migrate cơ sở dữ liệu: %v", err)
    }
    fmt.Println("Tạo bảng Users thành công!")

    var course models.Course
    err = db.AutoMigrate(&course)
    if err != nil {
        log.Fatalf("Không thể tự động migrate cơ sở dữ liệu: %v", err)
    }
    fmt.Println("Tạo bảng Course thành công!")
}