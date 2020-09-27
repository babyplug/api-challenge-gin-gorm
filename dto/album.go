package dto

type AlbumRequestForm struct {
	Name   string `json:"name" binding:"required"`
	Photos []uint `json:"photoId,omitempty"`
}
