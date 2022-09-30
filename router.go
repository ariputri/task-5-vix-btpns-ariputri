package main

import (
	"github.com/ariputri/btpn_api/controllers"
	"github.com/ariputri/btpn_api/database"
	"github.com/ariputri/btpn_api/middlewares"
	"github.com/ariputri/btpn_api/repository"
	"github.com/ariputri/btpn_api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                    = database.DatabaseConnection()
	userRepository  repository.UserRepository   = repository.NewUserRepository(db)
	photoRepository repository.PhotoRepository  = repository.NewPhotoRepository(db)
	jwtService      service.JWTService          = service.NewJWTService()
	userService     service.UserService         = service.NewUserService(userRepository)
	photoService    service.PhotoService        = service.NewPhotoService(photoRepository)
	authService     service.AuthService         = service.NewAuthService(userRepository)
	authController  controllers.AuthController  = controllers.NewAuthController(authService, jwtService)
	userController  controllers.UserController  = controllers.NewUserController(userService, jwtService)
	photoController controllers.PhotoController = controllers.NewPhotoController(photoService, jwtService)
)

func main() {
	defer database.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middlewares.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	photoRoutes := r.Group("api/photos", middlewares.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/", photoController.All)
		photoRoutes.POST("/", photoController.Insert)
		photoRoutes.GET("/:id", photoController.FindByID)
		photoRoutes.PUT("/:id", photoController.Update)
		photoRoutes.DELETE("/:id", photoController.Delete)
	}

	r.Run()
}
