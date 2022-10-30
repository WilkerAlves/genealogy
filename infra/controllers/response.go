package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseBadRequest(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": msg,
		"error":   err.Error(),
	})
}

func ResponseInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func ResponseNotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": msg,
	})
}
