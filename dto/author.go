package dto

type AuthorRequestform struct {
	Name   string `json:"name" binding:"required"`
	Photos []uint `json:"photoId,omitempty"`
}
