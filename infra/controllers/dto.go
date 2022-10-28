package controllers

type CreateOrUpdatePerson struct {
	Name string `json:"name" binding:"required,gte=1"`
}
