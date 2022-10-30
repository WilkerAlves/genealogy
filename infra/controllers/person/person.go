package person

import (
	"fmt"
	"net/http"

	"github.com/WilkerAlves/genealogy/application/person"
	"github.com/WilkerAlves/genealogy/infra/controllers"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	repo *repository.PersonRepository
}

func (ctrl *Controller) Create(c *gin.Context) {
	var body CreateOrUpdatePerson
	if err := c.BindJSON(&body); err != nil {
		controllers.ResponseBadRequest(c, "error while parser json", err)
		return
	}

	uc := person.NewCreatePersonUseCase(ctrl.repo)
	p, err := uc.Execute(c.Request.Context(), body.Name)
	if err != nil {
		controllers.ResponseInternalServerError(c, err)
		return
	}

	controllers.ResponseSuccess(c, p)
}

func (ctrl *Controller) FindById(c *gin.Context) {
	id, err := controllers.ValidateId(c.Param("id"))
	if err != nil {
		controllers.ResponseBadRequest(c, "", err)
		return
	}

	uc := person.NewFindPersonByIdUseCase(ctrl.repo)
	p, err := uc.Execute(c.Request.Context(), id)
	if err != nil {
		notFoundMsg := fmt.Sprintf("person %d not found", id)
		if err.Error() == notFoundMsg {
			controllers.ResponseNotFound(c, "person not found")
			return
		}

		controllers.ResponseInternalServerError(c, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	controllers.ResponseSuccess(c, p)
}

func (ctrl *Controller) Update(c *gin.Context) {
	id, err := controllers.ValidateId(c.Param("id"))
	if err != nil {
		controllers.ResponseBadRequest(c, "", err)
		return
	}

	var body CreateOrUpdatePerson
	if err := c.BindJSON(&body); err != nil {
		controllers.ResponseBadRequest(c, "error while parser json", err)
		return
	}

	uc := person.NewUpdatePersonUseCase(ctrl.repo)
	err = uc.Execute(c.Request.Context(), id, body.Name)
	if err != nil {
		notFoundMsg := fmt.Sprintf("person %d not found", id)
		if err.Error() == notFoundMsg {
			controllers.ResponseNotFound(c, notFoundMsg)
			return
		}

		controllers.ResponseInternalServerError(c, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	controllers.ResponseSuccess(c, nil)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id, err := controllers.ValidateId(c.Param("id"))
	if err != nil {
		controllers.ResponseBadRequest(c, "", err)
		return
	}

	uc := person.NewDeletePersonUseCase(ctrl.repo)
	err = uc.Execute(c.Request.Context(), id)
	if err != nil {
		notFoundMsg := fmt.Sprintf("person %d not found", id)
		if err.Error() == notFoundMsg {
			controllers.ResponseNotFound(c, notFoundMsg)
			return
		}

		controllers.ResponseInternalServerError(c, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	controllers.ResponseSuccess(c, nil)
}

func NewController(repo *repository.PersonRepository) *Controller {
	return &Controller{repo: repo}
}
