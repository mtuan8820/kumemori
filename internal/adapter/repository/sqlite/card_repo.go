package sqlite

import "gorm.io/gorm"

type CardRepo struct {
	DB *gorm.DB
}

func NewCardRepo(db *gorm.DB) *CardRepo {
	return &CardRepo{DB: db}
}
