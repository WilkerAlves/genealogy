package relationship

import (
	"github.com/WilkerAlves/genealogy/application/relationship"
	"github.com/WilkerAlves/genealogy/infra/controllers"
	"github.com/WilkerAlves/genealogy/infra/repository"
	"github.com/gin-gonic/gin"
)

//var repo *repository.RelationshipRepository

//func init() {
//	var err error
//	repo, err = repository.NewRelationshipRepository(os.Getenv("CONNECTION_STRING_DB"))
//	if err != nil {
//		log.Fatal().Err(err).Msg("error for create relationship repository")
//	}
//}

type Controller struct {
	repo *repository.RelationshipRepository
}

func (ctrl *Controller) Add(c *gin.Context) {
	var body CreateRelationship
	if err := c.BindJSON(&body); err != nil {
		controllers.ResponseBadRequest(c, "error while parser json", err)
		return
	}

	uc := relationship.NewAddUseCase(ctrl.repo)

	input := relationship.InputRelationship{
		Parent:   body.Parent,
		Children: body.Children,
	}

	err := uc.Execute(c.Request.Context(), input)
	if err != nil {
		controllers.ResponseInternalServerError(c, err)
		return
	}

	controllers.ResponseSuccess(c, nil)
}

func (ctrl *Controller) Genealogy(c *gin.Context) {
	id, err := controllers.ValidateId(c.Param("id"))
	if err != nil {
		controllers.ResponseBadRequest(c, "", err)
		return
	}

	uc := relationship.NewFindUseCase(ctrl.repo)
	result, err := uc.Execute(c.Request.Context(), id)
	if err != nil {
		controllers.ResponseInternalServerError(c, err)
		return
	}

	controllers.ResponseSuccess(c, result)
}

func (ctrl *Controller) Find(c *gin.Context) {
	id, err := controllers.ValidateId(c.Query("id"))
	if err != nil {
		controllers.ResponseBadRequest(c, "", err)
		return
	}

	findRelationship, err := controllers.ValidateId(c.Query("findrelationship"))
	if err != nil {
		controllers.ResponseBadRequest(c, "", err)
		return
	}

	uc := relationship.NewGetUseCase(ctrl.repo)
	result, err := uc.Execute(c.Request.Context(), id, findRelationship)
	if err != nil {
		controllers.ResponseInternalServerError(c, err)
		return
	}

	controllers.ResponseSuccess(c, result)
}

func NewController(repo *repository.RelationshipRepository) *Controller {
	return &Controller{repo: repo}
}
