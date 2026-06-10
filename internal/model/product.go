package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string
	Price     float64
	Stock     int
	CreatedAt time.Time
}
