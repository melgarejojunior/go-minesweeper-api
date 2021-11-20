package models

import (
	"time"

	"gorm.io/gorm"
)

type Field struct {
	ID          uint           `json:"-" gorm:"primaryKey"`
	Position    int            `json:"-"`
	Row         int            `json:"row"`
	Column      int            `json:"column"`
	IsBomb      bool           `json:"is_bomb"`
	BombsAround int            `json:"bombs_around"`
	IsOpened    bool           `json:"is_opened"`
	GameID      uint           `json:"-"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
