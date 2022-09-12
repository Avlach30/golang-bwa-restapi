package auth

import (
	"campaigns-restapi/helper"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *userHandler {
	return &userHandler{userService}
}

//* Mendeklarasikan fungsi untuk menghandle / mengurus input dari client dan response
func (handler *userHandler) SignUpHandler(context *gin.Context) {
	//* Handling multipart/form-data input type 

	input := SignUpInput{
		//* Handle multipart/form-data with input type plain text
		Name: context.PostForm("name"),
		Email: context.PostForm("email"),
		Occupation: context.PostForm("occupation"),
		Password: context.PostForm("password"),
	}

	err := context.ShouldBind(&input) //* Mengecek content-type dari input request client
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMsg := gin.H{"errors": errors}

		errorResponse := helper.ApiResponse(false, "Error occured", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	image, err := context.FormFile("image") //* Handle multipart/form-data with input type file
	if (err != nil) {
		errorResponse := helper.ApiResponse(false, "Error occured", err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	//* extract file extension
    extension := filepath.Ext(image.Filename)

	//* If file uploaded not image
	if (extension != ".jpg" && extension != ".jpeg" && extension != ".png") {
		errorResponse := helper.ApiResponse(false, "Error occured", "Sorry, only image file type can be uploaded")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	//* Uploaded limit less than 1,5 mb
	if (image.Size > 1572864) {
		errorResponse := helper.ApiResponse(false, "Error occured", "Sorry, image uploaded is more than limit (1,5 mb)")
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

    //* Generate random file name for the new uploaded file
    newFileName := uuid.New().String() + extension

	//* define uploaded file path
	imagePath := "images/" + newFileName

	//* Save uploaded file with desired destinations
	err = context.SaveUploadedFile(image, imagePath)
	if (err != nil) {
		errorResponse := helper.ApiResponse(false, "Error occured", err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	//* Update input struct
	input.AvatarFileName = imagePath


	//* Calling service for checking user availability from email
	isUserExist, err := handler.userService.CheckUserAvailabilityByEmail(input)
	if (isUserExist) {
		errorResponse := helper.ApiResponse(false, "Error occured", err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	//* Memanggil SignUp service untuk memproses hasil konversi input
	user, err := handler.userService.SignUp(input)
	if err != nil {
		errorResponse := helper.ApiResponse(false, "Error occured", err.Error())
		//* Mengembalikan JSON
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	
	responseFormatter := FormatUserSignupResponse(user, imagePath)
	successResponse := helper.ApiResponse(true, "Sign up successfully", responseFormatter)

	context.JSON(http.StatusCreated, successResponse)
}

func (handler *userHandler) LogInHandler(context *gin.Context) {
	var input LogInInput //* Mendefinisikan variabel dengan type LogInInput struct

	err := context.ShouldBindJSON(&input) //*Konversi JSON input ke struct
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMsg := gin.H{"errors": errors}

		errorResponse := helper.ApiResponse(false, "Error occured", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

	user, token, err := handler.userService.LogIn(input)
	if err != nil {
		errorResponse := helper.ApiResponse(false, "Error occured", err.Error())

		context.JSON(http.StatusUnauthorized, errorResponse)
		return
	}
	
	responseFormatter := FormatUserLoginResponse(user, token)
	successResponse := helper.ApiResponse(true, "Log in successfully", responseFormatter)

	context.JSON(http.StatusOK, successResponse)
}
