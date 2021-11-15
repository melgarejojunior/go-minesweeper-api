package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	Fields        *[]Field `json:"fields" gorm:"foreignKey:GameID;references:ID"`
	MinesweeperID uint     `json:"minesweeper_id"`
	Minesweeper   Minesweeper
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
