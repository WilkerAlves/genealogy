package relationship

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/relationship"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var repo *repository.RelationshipRepository

func init() {
	var err error
	repo, err = repository.NewRelationshipRepository(os.Getenv("CONNECTION_STRING_DB"))
	if err != nil {
		log.Fatal().Err(err).Msg("error for create relationship repository")
	}
}

func Add(c *gin.Context) {
	var body CreateRelationship
	if err := c.BindJSON(&body); err != nil {
		msg := fmt.Errorf("error while parser json: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg.Error(),
		})
	}

	uc := relationship.NewAddUseCase(repo)

	input := relationship.InputRelationship{
		Parent:   body.Parent,
		Children: body.Children,
	}

	err := uc.Execute(c.Request.Context(), input)
	if err != nil {
		msg := fmt.Errorf("error for create relationship: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func Genealogy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("invalid id: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg.Error(),
		})
		return
	}

	uc := relationship.NewFindUseCase(repo)
	result, err := uc.Execute(c.Request.Context(), id)
	if err != nil {
		msg := fmt.Errorf("error while get genealogy: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func Find(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		msg := fmt.Errorf("invalid id: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg.Error(),
		})
		return
	}

	findrelationship, err := strconv.Atoi(c.Query("findrelationship"))
	if err != nil {
		msg := fmt.Errorf("invalid id: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg.Error(),
		})
		return
	}

	uc := relationship.NewGetUseCase(repo)
	result, err := uc.Execute(c.Request.Context(), id, findrelationship)
	if err != nil {
		msg := fmt.Errorf("error while get genealogy: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
