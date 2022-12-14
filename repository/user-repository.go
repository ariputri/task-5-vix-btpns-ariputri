package repository

import (
	"log"

	"github.com/ariputri/btpn_api/app"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// a contract what UserRepository can do to db
type UserRepository interface {
	InsertUser(user app.User) app.User
	UpdateUser(user app.User) app.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) app.User
	ProfileUser(userID string) app.User
}

type userConnection struct {
	connection *gorm.DB
}

// creates new UserRepository instantly
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user app.User) app.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user app.User) app.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser app.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user app.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user app.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) app.User {
	var user app.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) app.User {
	var user app.User
	db.connection.Preload("Photos").Preload("Photos.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
