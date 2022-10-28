package person

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/WilkerAlves/genealogy/use_case/person"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

var repo *repository.PersonRepository

func init() {
	var err error
	repo, err = repository.NewPersonRepository(os.Getenv("CONNECTION_STRING_DB"))
	if err != nil {
		log.Fatal().Err(err).Msg("error for create person repository")
	}
}

func Create(c *gin.Context) {
	var body CreateOrUpdatePerson
	if err := c.BindJSON(&body); err != nil {
		msg := fmt.Errorf("error while parser json: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
	}

	uc := person.NewCreatePersonUseCase(repo)
	p, err := uc.Execute(c.Request.Context(), body.Name)
	if err != nil {
		msg := fmt.Errorf("error for create person: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   p.ID,
		"name": p.Name,
	})
}

func FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("invalid id: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
	}

	uc := person.NewFindPersonByIdUseCase(repo)
	p, err := uc.Execute(c.Request.Context(), id)
	if err != nil {
		notFoundMsg := fmt.Sprintf("person %d not found", id)
		if err.Error() == notFoundMsg {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFoundMsg,
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   p.ID,
		"name": p.Name,
	})
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("invalid id: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
	}

	var body CreateOrUpdatePerson
	if err := c.BindJSON(&body); err != nil {
		msg := fmt.Errorf("error while parser json: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
	}

	uc := person.NewUpdatePersonUseCase(repo)
	err = uc.Execute(c.Request.Context(), id, body.Name)
	if err != nil {
		notFoundMsg := fmt.Sprintf("person %d not found", id)
		if err.Error() == notFoundMsg {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFoundMsg,
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, nil)
}

func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Errorf("invalid id: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
	}

	uc := person.NewDeletePersonUseCase(repo)
	err = uc.Execute(c.Request.Context(), id)
	if err != nil {
		notFoundMsg := fmt.Sprintf("person %d not found", id)
		if err.Error() == notFoundMsg {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFoundMsg,
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, nil)
}
