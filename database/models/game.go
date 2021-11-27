package models

import (
	"minesweeper/database/models/status"
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	Fields        *[]Field          `json:"fields" gorm:"foreignKey:GameID;references:ID"`
	MinesweeperID uint              `json:"-"`
	Minesweeper   Minesweeper       `json:"minesweeper"`
	GameStatus    status.GameStatus `json:"game_status"`
	CreatedAt     time.Time         `json:"-"`
	UpdatedAt     time.Time         `json:"-"`
	DeletedAt     gorm.DeletedAt    `json:"-" gorm:"index"`
}
