package controllers

import (
	"errors"
	"fmt"
	"math/rand"
	"minesweeper/database"
	"minesweeper/database/models"
	"minesweeper/domain/entities"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Play(c *gin.Context) {
	response := Response{c}
	defer response.TreatError()

	gameID := response.Param("id")

	play := entities.Matrix{
		Row:    -1,
		Column: -1,
	}

	if err := response.BindJSON(&play); err != nil || play.Column < 0 || play.Row < 0 {
		if err != nil {
			panic(err)
		} else {
			panic(errors.New("Wrong values for Play"))
		}
	}

	var game models.Game
	if err := retrieveGame(&game, gameID, true); err != nil {
		panic(err)
	}

	if game.GameStatus == models.GameOver || game.GameStatus == models.Winner {
		panic(errors.New(fmt.Sprint("This game has already ended. Game status:", game.GameStatus)))
	}

	if game.GameStatus == models.NotStarted {
		if err := firstPlay(&game, play); err != nil {
			panic(err)
		}
	}

	if err := updatePlay(&game, play); err != nil {
		panic(err)
	}

	retrieveGame(&game, gameID, false)
	response.EmitSuccess(game)
}

func GetGame(c *gin.Context) {
	response := Response{c}
	defer response.TreatError()

	gameID := response.Param("id")

	var game models.Game
	if err := retrieveGame(&game, gameID, false); err != nil {
		panic(err)
	}

	response.EmitSuccess(game)
}

func firstPlay(game *models.Game, play entities.Matrix) error {
	var bombPositions []int
	totalPositions := game.Minesweeper.Column * game.Minesweeper.Row
	playPosition := play.Row*game.Minesweeper.Column + play.Column
	rand.Seed(time.Now().Unix())
	for len(bombPositions) != game.Minesweeper.NumOfBombs {
		bombPosition := rand.Intn(totalPositions)
		if bombPosition == playPosition || contains(bombPositions, bombPosition) {
			continue
		}
		bombPositions = append(bombPositions, bombPosition)
	}

	db := database.GetDatabase()

	for _, bomb := range bombPositions {
		if err := db.Model(models.Field{}).Where("position = ? and game_id = ?", bomb, game.ID).Update("is_bomb", true).Error; err != nil {
			return err
		}

		row := bomb / game.Minesweeper.Column
		column := bomb % game.Minesweeper.Column

		for i := row - 1; i <= row+1; i++ {
			for j := column - 1; j <= column+1; j++ {
				if i < 0 || i >= game.Minesweeper.Row || j < 0 || j >= game.Minesweeper.Column || (i == row && j == column) {
					continue
				}
				position := i*game.Minesweeper.Column + j
				var field models.Field
				if err := db.First(&field, "position = ? and game_id = ?", position, game.ID).Error; err != nil {
					return err
				}

				if field.IsBomb {
					continue
				}
				field.BombsAround = field.BombsAround + 1
				db.Save(&field)
			}
		}
	}

	updateGameStatus(game, models.Playing)

	return nil
}

func updatePlay(game *models.Game, play entities.Matrix) error {
	db := database.GetDatabase()

	var field models.Field
	err := db.Where("position = ? AND game_id = ?", (play.Row*game.Minesweeper.Column + play.Column), game.ID).First(&field).Error
	if err != nil {
		return err
	}

	if field.IsOpened {
		return errors.New("Field already opened")
	}
	field.IsOpened = true
	db.Save(field)

	if field.IsBomb {
		game.GameStatus = models.GameOver
		db.Save(game)
		return nil
	}

	retrieveGame(game, strconv.FormatUint(uint64(game.ID), 10), true)

	conditionPlayed := func(f models.Field) bool { return f.IsOpened }
	playedFields := len(filter(*game.Fields, conditionPlayed))
	totalPos := game.Minesweeper.Column * game.Minesweeper.Row

	if playedFields == (totalPos - game.Minesweeper.NumOfBombs) {
		updateGameStatus(game, models.Winner)
	}

	return nil
}

func updateGameStatus(game *models.Game, status models.GameStatus) {
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

func retrieveGame(game *models.Game, gameID string, allFields bool) error {
	db := database.GetDatabase()

	err := db.First(&game, gameID).Error
	if err != nil {
		return err
	}

	fields := []models.Field{}
	if allFields || game.GameStatus == models.GameOver || game.GameStatus == models.Winner {
		if err := db.Order("position").Find(&fields, "game_id = ?", game.ID).Error; err != nil {
			return err
		}
	} else {
		if err := db.Order("position").Find(&fields, "game_id = ? and is_opened = true", game.ID).Error; err != nil {
			return err
		}
	}
	game.Fields = &fields

	var ms models.Minesweeper
	if err := db.Find(&ms, game.MinesweeperID).Error; err != nil {
		return err
	}
	game.Minesweeper = ms

	return nil
}
