package repository

import (
	"log"

	"kumemori/internal/domain/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDb opens a SQLite connection and auto-migrates the Card and Deck models.
// Returns the DB instance or an error if migration fails.
func InitDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.Card{}, &model.Deck{})
	if err != nil {
		return nil, err
	}

	log.Println("Db migrated successfully")
	return db, nil
}
