package auth

import (
	"campaigns-restapi/helper"
	"net/http"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService Service
}

func NewUserHandler(userService Service) *userHandler {
	return &userHandler{userService}
}

//* Mendeklarasikan fungsi untuk menghandle / mengurus input dari client dan response
func (handler *userHandler) SignUpHandler(context *gin.Context) {
	var input SignUpInput //* Mendefinisikan variabel dengan type SignUpInput struct

	//*Konversi JSON input ke struct
	err := context.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		//* Mapping ke object
		errorMsg := gin.H{"errors": errors}

		errorResponse := helper.ApiResponse(false, "Error occured", errorMsg)
		context.JSON(http.StatusUnprocessableEntity, errorResponse)
		return
	}

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
	
	responseFormatter := FormatUserSignupResponse(user)
	successResponse := helper.ApiResponse(true, "Sign up successfully", responseFormatter)

	context.JSON(http.StatusCreated, successResponse)
}

func (handler *userHandler) LogInHandler(context *gin.Context) {
	var input LogInInput

	err := context.ShouldBindJSON(&input)
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
