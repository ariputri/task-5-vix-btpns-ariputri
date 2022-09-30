package repository

import (
	"github.com/ariputri/btpn_api/app"
	"gorm.io/gorm"
)

// PhotoRepository is a ....
type PhotoRepository interface {
	InsertPhoto(b app.Photo) app.Photo
	UpdatePhoto(b app.Photo) app.Photo
	DeletePhoto(b app.Photo)
	AllPhoto() []app.Photo
	FindPhotoByID(photoID uint64) app.Photo
}

type photoConnection struct {
	connection *gorm.DB
}

// NewPhotoRepository creates an instance PhotoRepository
func NewPhotoRepository(dbConn *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: dbConn,
	}
}

func (db *photoConnection) InsertPhoto(b app.Photo) app.Photo {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *photoConnection) UpdatePhoto(b app.Photo) app.Photo {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *photoConnection) DeletePhoto(b app.Photo) {
	db.connection.Delete(&b)
}

func (db *photoConnection) FindPhotoByID(photoID uint64) app.Photo {
	var photo app.Photo
	db.connection.Preload("User").Find(&photo, photoID)
	return photo
}

func (db *photoConnection) AllPhoto() []app.Photo {
	var photos []app.Photo
	db.connection.Preload("User").Find(&photos)
	return photos
}
