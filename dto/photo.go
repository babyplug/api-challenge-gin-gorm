package dto

type PhotoRequestform struct {
	Description string `json:"description" binding:"required"`
	FileName    string `json:"fileName" binding:"required"`
	IsPublished bool   `json:"isPublished" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Views       int64  `json:"views" binding:"required"`
	AuthorId    uint   `json:"authorId" binding:"required"`
}
