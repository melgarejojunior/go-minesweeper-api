package domain

import (
	"errors"
	"fmt"
	"minesweeper/database"
	"minesweeper/database/models"
	"minesweeper/domain/entities"
)

func CreateNewMinesweeper(cMinesweeper entities.CreateMinesweeper) *models.Game {

	if cMinesweeper.NumOfBombs >= (cMinesweeper.Column * cMinesweeper.Row) {
		errstr := fmt.Sprint("Number of bombs(", cMinesweeper.NumOfBombs,
			") cannot be greater or equals column (",
			cMinesweeper.Column, ") * row(", cMinesweeper.Row, ") = ",
			(cMinesweeper.Column * cMinesweeper.Row))
		panic(errors.New(errstr))
	}

	db := database.GetDatabase()

	newMinesweeper := models.Minesweeper{
		Row:        cMinesweeper.Row,
		Column:     cMinesweeper.Column,
		NumOfBombs: cMinesweeper.NumOfBombs,
	}

	if err := db.Create(&newMinesweeper).Error; err != nil {
		panic(err)
	}
	game := models.Game{
		ID:            0,
		Fields:        &[]models.Field{},
		MinesweeperID: newMinesweeper.ID,
		Minesweeper:   newMinesweeper,
		GameStatus:    models.NotStarted,
	}

	if err := db.Create(&game).Error; err != nil {
		panic(err)
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
				IsOpened:    false,
				BombsAround: 0,
				GameID:      game.ID,
			}

			if err := db.Create(&field).Error; err != nil {
				panic(err)
			}

			fields = append(fields, field)
			position++
		}
	}

	game.Fields = &fields

	if err := db.Save(&game).Error; err != nil {
		panic(err)
	}

	return &game
}
