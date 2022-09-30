package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ariputri/btpn_api/app"
	"github.com/ariputri/btpn_api/helpers"
	"github.com/ariputri/btpn_api/models"
	"github.com/ariputri/btpn_api/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// PhotoController is a ...
type PhotoController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type photoController struct {
	photoService service.PhotoService
	jwtService   service.JWTService
}

// NewPhotoController create a new instances of PhotoController
func NewPhotoController(photoServ service.PhotoService, jwtServ service.JWTService) PhotoController {
	return &photoController{
		photoService: photoServ,
		jwtService:   jwtServ,
	}
}

func (c *photoController) All(context *gin.Context) {
	var photos []app.Photo = c.photoService.All()
	res := helpers.BuildResponse(true, "OK", photos)
	context.JSON(http.StatusOK, res)
}

func (c *photoController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helpers.BuildErrorResponse("No param id was found", err.Error(), helpers.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var photo app.Photo = c.photoService.FindByID(id)
	if (photo == app.Photo{}) {
		res := helpers.BuildErrorResponse("Data not found", "No data with given id", helpers.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helpers.BuildResponse(true, "OK", photo)
		context.JSON(http.StatusOK, res)
	}
}

func (c *photoController) Insert(context *gin.Context) {
	var photoCreateModel models.PhotoCreateModel
	errModel := context.ShouldBind(&photoCreateModel)
	if errModel != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errModel.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			photoCreateModel.UserID = convertedUserID
		}
		result := c.photoService.Insert(photoCreateModel)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *photoController) Update(context *gin.Context) {
	var photoUpdateModel models.PhotoUpdateModel
	errModel := context.ShouldBind(&photoUpdateModel)
	if errModel != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errModel.Error(), helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoService.IsAllowedToEdit(userID, photoUpdateModel.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			photoUpdateModel.UserID = id
		}
		result := c.photoService.Update(photoUpdateModel)
		response := helpers.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) Delete(context *gin.Context) {
	var photo app.Photo
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helpers.BuildErrorResponse("Failed tou get id", "No param id were found", helpers.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	photo.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.photoService.IsAllowedToEdit(userID, photo.ID) {
		c.photoService.Delete(photo)
		res := helpers.BuildResponse(true, "Deleted", helpers.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *photoController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
