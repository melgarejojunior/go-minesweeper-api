package migrations

import (
	"minesweeper/database/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.Minesweeper{}, &models.Game{}, &models.Field{})
}
