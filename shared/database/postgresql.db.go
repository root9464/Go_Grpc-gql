package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

var DB Database

func ConnectDb() (*Database, error) {
	_ = godotenv.Load()

	// dbURL, exists := os.LookupEnv("PSQL_URL")
	// if !exists {
	// 	log.Fatal("PSQL_URL is not set in environment variables")
	// }

	dbURL := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(200)

	DB = Database{Db: db}
	return &DB, nil
}
