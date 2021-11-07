package migrations

import (
	"minesweeper/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Minesweeper{})
}
