package app

// Photo struct represents photos table in database
type Photo struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:text" json:"caption"`
	PhotoUrl string `gorm:"size:254;" json:"photo_url"`
	UserID   uint64 `gorm:"not null" json:"-"`
	User     User   `gorm:"foreignkey:UserID" json:"user_id" binding:"required"`
}
