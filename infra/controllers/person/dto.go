package person

type CreateOrUpdatePerson struct {
	Name string `json:"name" binding:"required"`
}
