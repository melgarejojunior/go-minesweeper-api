package domain

import (
	"minesweeper/database"
	"minesweeper/database/models"
	"minesweeper/domain/entities"
)

func Execute(cMinesweeper entities.CreateMinesweeper, onFailure func(error)) *models.Minesweeper {
	db := database.GetDatabase()

	newMinesweeper := models.Minesweeper{
		Row:        cMinesweeper.Row,
		Column:     cMinesweeper.Column,
		NumOfBombs: cMinesweeper.NumOfBombs,
	}

	err := db.Create(&newMinesweeper).Error
	if err != nil {
		onFailure(err)
		return nil
	}
	return &newMinesweeper
}
