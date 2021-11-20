package controllers

import (
	"minesweeper/domain"
	"minesweeper/domain/entities"

	"github.com/gin-gonic/gin"
)

func ConfigMinesweeper(c *gin.Context) {
	response := Response{c}
	defer response.TreatError()

	var createMinesweeper entities.CreateMinesweeper

	if err := response.ShouldBindJSON(&createMinesweeper); err != nil {
		panic(err)
	}

	newMinesweeper := domain.CreateNewMinesweeper(createMinesweeper)

	if newMinesweeper != nil {
		response.EmitSuccess(newMinesweeper)
	}
}
