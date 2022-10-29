package relationship

type CreateRelationship struct {
	Parent   int `json:"parent" binding:"required,gte=1"`
	Children int `json:"children" binding:"required,gte=1"`
}
