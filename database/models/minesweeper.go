package models

import (
	"time"

	"gorm.io/gorm"
)

type Minesweeper struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Row        int            `json:"row"`
	Column     int            `json:"column"`
	NumOfBombs int            `json:"num_of_bombs"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
