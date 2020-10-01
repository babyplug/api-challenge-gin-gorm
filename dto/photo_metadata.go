package dto

type PhotoMetadataRequestform struct {
	Height      int64  `json:"height" binding:"required"`
	Width       int64  `json:"width" binding:"required"`
	Orientation string `json:"orientation" binding:"required"`
	Compressed  int64  `json:"compressed" binding:"required"`
	Comment     string `json:"comment" binding:"required"`
	PhotoId     uint   `json:"photoId" binding:"required"`
}
