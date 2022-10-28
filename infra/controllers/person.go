package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePerson(c *gin.Context) {
	var body CreateOrUpdatePerson
	if err := c.BindJSON(&body); err != nil {
		msg := fmt.Errorf("error while parser json: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "chegue aqui CreatePerson",
	})
}

func GetPerson(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"message": "chegue aqui GetPerson",
		"id":      id,
	})
}

func GetPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "chegue aqui GetPersons",
		"persons": make([]string, 0),
	})
}

func UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var body CreateOrUpdatePerson
	if err := c.BindJSON(&body); err != nil {
		msg := fmt.Errorf("error while parser json: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
			"id":      id,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "chegue aqui UpdatePerson",
	})
}

func DeletePerson(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "chegue aqui DeletePerson",
		"id":      id,
	})
}
