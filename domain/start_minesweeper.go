package domain

import (
	"errors"
	"minesweeper/database"
	"minesweeper/database/models"
	"minesweeper/domain/entities"
)

func Execute(cMinesweeper entities.CreateMinesweeper, onFailure func(error)) *models.Game {

	if cMinesweeper.NumOfBombs >= (cMinesweeper.Column * cMinesweeper.Row) {
		err := errors.New("Number of bombs cannot be greater or equals column * row")
		onFailure(err)
		return nil
	}

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
	game := models.Game{
		MinesweeperID: newMinesweeper.ID,
		Minesweeper:   newMinesweeper,
	}

	err = db.Create(&game).Error
	if err != nil {
		onFailure(err)
		return nil
	}

	fields := []models.Field{}

	position := 0
	for row := 0; row < newMinesweeper.Row; row++ {
		for column := 0; column < newMinesweeper.Column; column++ {
			field := models.Field{
				Position:    position,
				Row:         row,
				Column:      column,
				IsBomb:      false,
				BombsAround: 0,
				GameID:      game.ID,
			}

			err = db.Create(&field).Error
			if err != nil {
				onFailure(err)
				return nil
			}

			fields = append(fields, field)
			position++
		}
	}

	game.Fields = &fields
	err = db.Save(&game).Error
	if err != nil {
		onFailure(err)
		return nil
	}

	return &game
}
