package domain

import (
	"errors"
	"fmt"
	"math/rand"
	"minesweeper/database"
	"minesweeper/database/models"
	"minesweeper/database/models/status"
	"minesweeper/domain/entities"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func MakeAPlay(gameID string, play entities.Matrix) models.Game {
	var game models.Game
	if err := RetrieveGame(&game, gameID, true); err != nil {
		panic(err)
	}

	if game.GameStatus == status.GameOver || game.GameStatus == status.Winner {
		panic(errors.New(fmt.Sprint("This game has already ended. Game status:", game.GameStatus.ConvertToValue())))
	}

	if game.GameStatus == status.NotStarted {
		if err := firstPlay(&game, play); err != nil {
			panic(err)
		}
	}

	if err := updatePlay(&game, play); err != nil {
		panic(err)
	}

	if err := RetrieveGame(&game, gameID, false); err != nil {
		panic(err)
	}
	return game
}

func firstPlay(game *models.Game, play entities.Matrix) error {
	var bombPositions []int
	totalPositions := game.Minesweeper.Column * game.Minesweeper.Row
	playPosition := getPosition(play, game.Minesweeper)
	bombPositions = generateBombPositions(bombPositions, game, totalPositions, playPosition)

	populateBombs(bombPositions, game)

	updateGameStatus(game, status.Playing)

	return nil
}

func generateBombPositions(bombPositions []int, game *models.Game, totalPositions int, playPosition int) []int {
	rand.Seed(time.Now().Unix())
	for len(bombPositions) != game.Minesweeper.NumOfBombs {
		bombPosition := rand.Intn(totalPositions)
		if bombPosition == playPosition || contains(bombPositions, bombPosition) {
			continue
		}
		bombPositions = append(bombPositions, bombPosition)
	}
	return bombPositions
}

func populateBombs(bombPositions []int, game *models.Game) {
	db := database.GetDatabase()

	for _, bomb := range bombPositions {
		if err := db.Model(models.Field{}).Where("position = ? and game_id = ?", bomb, game.ID).Update("is_bomb", true).Error; err != nil {
			panic(err)
		}

		row := bomb / game.Minesweeper.Column
		column := bomb % game.Minesweeper.Column

		setNumOfBombsAround(row, column, game, db)
	}
}

func setNumOfBombsAround(row int, column int, game *models.Game, db *gorm.DB) {
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i < 0 || i >= game.Minesweeper.Row || j < 0 || j >= game.Minesweeper.Column || (i == row && j == column) {
				continue
			}
			position := i*game.Minesweeper.Column + j
			var field models.Field
			if err := db.First(&field, "position = ? and game_id = ?", position, game.ID).Error; err != nil {
				panic(err)
			}

			if field.IsBomb {
				continue
			}
			field.BombsAround = field.BombsAround + 1
			db.Save(&field)
		}
	}
}

func updatePlay(game *models.Game, play entities.Matrix) error {

	if shouldReturn, returnValue := updateField(play, game); shouldReturn {
		return returnValue
	}

	if err := RetrieveGame(game, strconv.FormatUint(uint64(game.ID), 10), true); err != nil {
		return err
	}

	conditionPlayed := func(f models.Field) bool { return f.IsOpened }
	playedFields := len(filter(*game.Fields, conditionPlayed))
	totalPos := game.Minesweeper.Column * game.Minesweeper.Row

	if playedFields == (totalPos - game.Minesweeper.NumOfBombs) {
		updateGameStatus(game, status.Winner)
	}

	return nil
}

func updateField(play entities.Matrix, game *models.Game) (bool, error) {
	db := database.GetDatabase()

	var field models.Field

	if err := db.Where("position = ? AND game_id = ?", getPosition(play, game.Minesweeper), game.ID).First(&field).Error; err != nil {
		return true, err
	}

	if field.IsOpened {
		return true, errors.New("Field already opened")
	}
	field.IsOpened = true
	db.Save(field)

	if field.IsBomb {
		updateGameStatus(game, status.GameOver)
		return true, nil
	}
	return false, nil
}

func updateGameStatus(game *models.Game, status status.GameStatus) {
	db := database.GetDatabase()
	game.GameStatus = status
	db.Save(game)
}

func filter(fields []models.Field, conditionFunc func(models.Field) bool) (ret []models.Field) {
	for _, f := range fields {
		if conditionFunc(f) {
			ret = append(ret, f)
		}
	}
	return
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func getPosition(play entities.Matrix, minesweeper models.Minesweeper) int {
	return play.Row*minesweeper.Column + play.Column
}
