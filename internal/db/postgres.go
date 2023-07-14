package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitPostgreSQL initializes the PostgreSQL database connection
func InitPostgreSQL() {
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")
	database := os.Getenv("PG_DATABASE")
	port := os.Getenv("PG_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable TimeZone=America/Sao_Paulo"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Postgres Connected!")

	// AutoMigrate models here if needed
	// db.AutoMigrate(&entity.User{}, &entity.Product{})
}

// GetDB returns the reference to the database connection
func GetDB() *gorm.DB {
	return db
}

// ClosePostgreSQL closes the PostgreSQL database connection
func ClosePostgreSQL() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			log.Println(err)
		} else {
			err := sqlDB.Close()
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Postgres Disconnected!")
}
