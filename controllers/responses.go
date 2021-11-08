package controllers

import "github.com/gin-gonic/gin"

type Response struct {
	*gin.Context
}

func (c Response) EmitError(code int, errorStr string) {
	c.JSON(code, gin.H{
		"error": errorStr,
	})
}

// func (c response) EmitError(code int, errorStr string, errors ...error) {
// 	c.JSON(code, gin.H{
// 		"error": errorStr,
// 	})
// }

func (c Response) EmitSuccess(a interface{}) {
	c.JSON(200, a)
}
