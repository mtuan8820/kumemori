package repository

import (
	"kumemori/internal/core/domain"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQliteDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDb() (*gorm.DB, error) {
	db, err := ConnectSQliteDb()
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sqlDB.Close()

	err = db.AutoMigrate(&domain.Card{})
	if err != nil {
		return nil, err
	}

	log.Println("Db migrated successfully")
	return db, nil
}
