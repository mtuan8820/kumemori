package entity

import "time"

type SQLiteDeck struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:text;not null;uniqueIndex"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	NewCardLimit  int `gorm:"default:20"`
	ReviewLimit   int `gorm:"default:100"`
	LastStudiedAt time.Time

	Cards []SQLiteCard `gorm:"constraint:OnDelete:CASCADE;"`
}
