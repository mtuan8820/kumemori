package sqlite

import (
	"kumemori/internal/domain/model"
	"log"

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

	// Ensure SQLite enforces foreign keys for cascading
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return nil, err
	}

	// Migrate parent before child so FK can be created
	err = db.AutoMigrate(&model.Deck{}, &model.Card{})
	if err != nil {
		return nil, err
	}

	log.Println("Db migrated successfully")
	return db, nil
}
