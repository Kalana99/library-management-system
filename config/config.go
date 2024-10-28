package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
	"fmt"
)

var DB *gorm.DB

func InitDB(host, port, user, password, dbname string) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	host, port, user, password, dbname)
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    log.Println("Database connected!")
}
