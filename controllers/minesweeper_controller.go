package controllers

import (
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
}
