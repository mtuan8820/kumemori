package repository

import (
	"kumemori/internal/core/domain/entity"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDb(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to in-memory sqlite: %v", err)
	}

	if err := db.AutoMigrate(&entity.Deck{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	return db
}

func TestCreateAndRead(t *testing.T) {
	db := setupTestDb(t)
	repo := NewDeckSqliteRepository(db)

	deck := entity.Deck{Name: "Test Deck"}
	if err := repo.CreateDeck(deck); err != nil {
		t.Fatalf("fail to create deck: %v", err)
	}

	got, err := repo.ReadDeck(1)
	if err != nil {
		t.Fatalf("failed to get deck: %v", err)
	}

	if got.Name != "Test Deck" {
		t.Errorf("got Name %q, want %q", got.Name, "Test Deck")
	}
}

func TestCreateAndDelete(t *testing.T) {}
