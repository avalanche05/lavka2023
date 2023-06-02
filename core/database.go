package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const dsn = "postgresql://postgres:password@db/postgres"

func GetDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Can't connect to DB, %v\n", err)
		os.Exit(1)
	}

	return db
}
