package models

import (
	"time"

	"gorm.io/gorm"
)

type Field struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Position    int            `json:"position"`
	Row         int            `json:"row"`
	Column      int            `json:"column"`
	IsBomb      bool           `json:"is_bomb"`
	BombsAround int            `json:"bombs_around"`
	IsOpened    bool           `json:"is_opened"`
	GameID      uint           `json:"game_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
