package controllers

import (
	"fmt"
	"minesweeper/domain"
	"minesweeper/domain/entities"

	"github.com/gin-gonic/gin"
)

func ConfigMinesweeper(c *gin.Context) {
	response := Response{c}

	var createMinesweeper entities.CreateMinesweeper

	err := response.ShouldBindJSON(&createMinesweeper)
	if err != nil {
		response.EmitError(400, err.Error())
		return
	}

	errCallback := func(e error) {
		fmt.Println(e.Error())
		response.EmitError(422, e.Error())
	}

	newMinesweeper := domain.Execute(createMinesweeper, errCallback)

	if newMinesweeper != nil {
		response.EmitSuccess(newMinesweeper)
	}
}
