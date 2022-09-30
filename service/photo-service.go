package service

import (
	"fmt"
	"log"

	"github.com/ariputri/btpn_api/app"
	"github.com/ariputri/btpn_api/models"
	"github.com/ariputri/btpn_api/repository"
	"github.com/mashingan/smapping"
)

// PhotoService is a ....
type PhotoService interface {
	Insert(b models.PhotoCreateModel) app.Photo
	Update(b models.PhotoUpdateModel) app.Photo
	Delete(b app.Photo)
	All() []app.Photo
	FindByID(PhotoID uint64) app.Photo
	IsAllowedToEdit(userID string, PhotoID uint64) bool
}

type photoService struct {
	photoRepository repository.PhotoRepository
}

// NewPhotoService .....
func NewPhotoService(photoRepo repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: photoRepo,
	}
}

func (service *photoService) Insert(b models.PhotoCreateModel) app.Photo {
	photo := app.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.photoRepository.InsertPhoto(photo)
	return res
}

func (service *photoService) Update(b models.PhotoUpdateModel) app.Photo {
	photo := app.Photo{}
	err := smapping.FillStruct(&photo, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.photoRepository.UpdatePhoto(photo)
	return res
}

func (service *photoService) Delete(b app.Photo) {
	service.photoRepository.DeletePhoto(b)
}

func (service *photoService) All() []app.Photo {
	return service.photoRepository.AllPhoto()
}

func (service *photoService) FindByID(photoID uint64) app.Photo {
	return service.photoRepository.FindPhotoByID(photoID)
}

func (service *photoService) IsAllowedToEdit(userID string, photoID uint64) bool {
	b := service.photoRepository.FindPhotoByID(photoID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
