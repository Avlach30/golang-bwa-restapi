package campaign

import (
	"campaigns-restapi/helper"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service Service
}

func NewCampaignHandler(service Service) *campaignHandler {
	return &campaignHandler{service}
}

func (handler *campaignHandler) GetCampaigns(context *gin.Context) {
	//* Get value from request query user_id and convert it to int
	userId, _ := strconv.Atoi(context.Query("user_id")) 

	campaigns, err := handler.service.FindAllCampaigns(userId)
	if (err != nil) {
		errorResponse := helper.ApiResponse(false, "Error occured", err.Error())
		context.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	response := helper.ApiResponse(true, "Get all campaigns successfully", FormatGetCampaignsResponse(campaigns))
	context.JSON(http.StatusOK, response)
}