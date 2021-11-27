package domain

import (
	"minesweeper/database"
	"minesweeper/database/models"
	"minesweeper/database/models/status"
	"sync"

	"gorm.io/gorm"
)

func RetrieveGame(game *models.Game, gameID string, allFields bool) error {
	db := database.GetDatabase()

	if err := db.First(&game, gameID).Error; err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		retrieveFields(allFields, game, db)
		wg.Done()
	}()
	go func() {
		retrieveMinesweeper(db, game)
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func retrieveMinesweeper(db *gorm.DB, game *models.Game) {
	var ms models.Minesweeper
	if err := db.Find(&ms, game.MinesweeperID).Error; err != nil {
		panic(err)
	}
	game.Minesweeper = ms
}

func retrieveFields(allFields bool, game *models.Game, db *gorm.DB) {
	fields := []models.Field{}

	if allFields || game.GameStatus == status.GameOver || game.GameStatus == status.Winner {
		if err := db.Order("position").Find(&fields, "game_id = ?", game.ID).Error; err != nil {
			panic(err)
		}
	} else {
		if err := db.Order("position").Find(&fields, "game_id = ? and is_opened = true", game.ID).Error; err != nil {
			panic(err)
		}
	}
	game.Fields = &fields
}