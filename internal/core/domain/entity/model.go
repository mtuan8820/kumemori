package entity

import "time"

type Card struct {
	ID        uint   `gorm:"primaryKey"`
	DeckID    uint   `gorm:"not null;index"`
	Front     string `gorm:"type:text;not null"`
	Back      string `gorm:"type:text;not null"`
	CreatedAt time.Time

	Repetitions  int       `gorm:"not null;default:0"`
	Lapses       int       `gorm:"not null;default:0"`
	EaseFactor   float64   `gorm:"not null;default:2.5"`
	Interval     int       `gorm:"not null;default:1"`
	Due          time.Time `gorm:"not null"`
	LastReviewed time.Time `gorm:"not null"`
}

type Deck struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:text;not null;uniqueIndex"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	NewCardLimit  int `gorm:"default:20"`
	ReviewLimit   int `gorm:"default:100"`
	LastStudiedAt time.Time

	Cards []Card `gorm:"constraint:OnDelete:CASCADE;"`
}
