package models

// PhotoUpdate is used for the user to update the photo
type PhotoUpdateModel struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" gorm:"size:254;" binding:"required"`
	UserID   uint64 `json:"user_id" form:"user_id"`
}

// PhotoCreate is used to upload a new photo
type PhotoCreateModel struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" gorm:"size:254;" binding:"required"`
	UserID   uint64 `json:"user_id" form:"user_id"`
}
