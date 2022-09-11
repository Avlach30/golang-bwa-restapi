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
