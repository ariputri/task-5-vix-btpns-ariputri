package service

import (
	"log"

	"github.com/ariputri/btpn_api/app"
	"github.com/ariputri/btpn_api/models"
	"github.com/ariputri/btpn_api/repository"
	"github.com/mashingan/smapping"
)

// UserService is a contract.....
type UserService interface {
	Update(user models.UserUpdateModel) app.User
	Profile(userID string) app.User
}

type userService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user models.UserUpdateModel) app.User {
	userToUpdate := app.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) app.User {
	return service.userRepository.ProfileUser(userID)
}
