package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB() {
    dsn := os.Getenv("DATABASE_URL")
    // dbname := os.Getenv("DATABASE_NAME")
    // createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbname)
    // DB.Exec(createDatabaseCommand)
    // DB.Exec("use %s", dbname)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    DB = db
}
