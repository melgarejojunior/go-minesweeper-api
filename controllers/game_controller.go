package controllers

import (
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

	gameID := response.Param("id")

	var play entities.Matrix

	err := response.ShouldBindJSON(&play)
	if err != nil {
		response.EmitError(422, err.Error())
		return
	}

	var game models.Game
	retrieveGame(&game, gameID)

	conditionUnplayed := func(f models.Field) bool { return !f.IsOpened }
	isTheFirstPlay := filter(*game.Fields, conditionUnplayed)

	if len(isTheFirstPlay) == len(*game.Fields) {
		err = firstPlay(&game, play)
		if err != nil {
			response.EmitError(422, err.Error())
			return
		}
	}

	err = updatePlay(&game, play)
	if err != nil {
		response.EmitError(422, err.Error())
		return
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
		db.Model(models.Field{}).Where("position = ? and game_id = ?", bomb, game.ID).Update("is_bomb", true)

		row := bomb / game.Minesweeper.Column
		column := bomb % game.Minesweeper.Column

		for i := row - 1; i <= row+1; i++ {
			for j := column - 1; j <= column+1; j++ {
				if i < 0 || i >= game.Minesweeper.Row || j < 0 || j >= game.Minesweeper.Column || (i == row && j == column) {
					continue
				}
				position := i*game.Minesweeper.Column + j
				var field models.Field
				db.First(&field, "position = ? and game_id = ?", position, game.ID)

				if field.IsBomb {
					continue
				}
				field.BombsAround = field.BombsAround + 1
				db.Save(&field)
			}
		}
	}

	retrieveGame(game, strconv.FormatUint(uint64(game.ID), 10))
	return nil
}

func updatePlay(game *models.Game, play entities.Matrix) error {
	db := database.GetDatabase()

	var field models.Field
	err := db.Where("position = ? AND game_id = ?", (play.Row*game.Minesweeper.Column + play.Column), game.ID).First(&field).Error
	if err != nil {
		return err
	}

	field.IsOpened = true
	db.Save(field)
	retrieveGame(game, strconv.FormatUint(uint64(game.ID), 10))

	return nil
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

func retrieveGame(game *models.Game, gameID string) error {
	db := database.GetDatabase()

	err := db.First(&game, gameID).Error
	if err != nil {
		return err
	}

	fields := []models.Field{}
	err = db.Find(&fields, "game_id = ?", game.ID).Error
	if err != nil {
		return err
	}
	game.Fields = &fields

	var ms models.Minesweeper
	err = db.Find(&ms, game.MinesweeperID).Error
	if err != nil {
		return err
	}
	game.Minesweeper = ms

	return nil
}
