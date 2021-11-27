package controllers

import (
	"errors"
	"minesweeper/database/models"
	"minesweeper/domain"
	"minesweeper/domain/entities"

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

	game := domain.MakeAPlay(gameID, play)

	response.EmitSuccess(game)
}

func GetGame(c *gin.Context) {
	response := Response{c}
	defer response.TreatError()

	gameID := response.Param("id")

	var game models.Game
	if err := domain.RetrieveGame(&game, gameID, false); err != nil {
		panic(err)
	}

	response.EmitSuccess(game)
}
