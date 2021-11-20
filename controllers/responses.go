package controllers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	*gin.Context
}

func (response Response) TreatError() {
	if r := recover(); r != nil {
		response.emitError(400, errors.New(fmt.Sprint(r)))
	}
}

func (response Response) emitError(code int, e error) {
	response.JSON(code, gin.H{
		"error": e.Error(),
	})
}

func (response Response) EmitSuccess(a interface{}) {
	response.JSON(200, a)
}
