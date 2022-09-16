package campaign

import "strings"

type GetCampaignformatterResponse struct {
	ID               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	Slug             string `json:"slug"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
}

type GetSpecifiedUserCampaignFormatterResponse struct {
	Name           string `json:"name"`
	AvatarFileName string `json:"avatar_file_name"`
}

type GetSpecifiedImageCampaignFormatterResponse struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type GetSpecifiedCampaignFormatterResponse struct {
	ID               int                                          `json:"id"`
	Name             string                                       `json:"name"`
	ShortDescription string                                       `json:"short_description"`
	ImageUrl         string                                       `json:"image_url"`
	GoalAmount       int                                          `json:"goal_amount"`
	CurrentAmount    int                                          `json:"current_amount"`
	UserId           int                                          `json:"user_id"`
	Description      string                                       `json:"description"`
	Perks            []string                                     `json:"perks"`
	User             GetSpecifiedUserCampaignFormatterResponse    `json:"user"`
	Images           []GetSpecifiedImageCampaignFormatterResponse `json:"images"`
}

//* format response for get each campaign
func FormatGetCampaignResponse(campaign Campaign) GetCampaignformatterResponse {
	format := GetCampaignformatterResponse{
		ID:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		Slug:             campaign.Slug,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
	}

	if len(campaign.CampaignImages) > 0 { //* Get imageurl from first index of filename
		format.ImageUrl = "/" + campaign.CampaignImages[0].FileName
	}

	return format
}

//* format response for get all campaigns with call function returned format each campaign
func FormatGetCampaignsResponse(campaigns []Campaign) []GetCampaignformatterResponse {

	//* Declarating campaigns formatter variable with intial value is empty array
	campaignsFormatter := []GetCampaignformatterResponse{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatGetCampaignResponse(campaign)           //* Call function for format get each campaign
		campaignsFormatter = append(campaignsFormatter, campaignFormatter) //* Append called function to array
	}

	return campaignsFormatter
}

func FormatGetSpecifiedCampaignResponse(campaign Campaign) GetSpecifiedCampaignFormatterResponse {
	format := GetSpecifiedCampaignFormatterResponse{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Description:      campaign.Description,
	}

	if len(campaign.CampaignImages) > 0 { //* Get imageurl from first index of filename
		format.ImageUrl = "/" + campaign.CampaignImages[0].FileName
	}

	//* convert campaign perk for array splitted by comma
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") { 
		perks = append(perks, perk)
	}

	format.Perks = perks

	//* get campaign owner for define a user campaign
	campaignUser := campaign.User
	formatCampaignUser := GetSpecifiedUserCampaignFormatterResponse{
		Name:           campaignUser.Name,
		AvatarFileName: campaignUser.AvatarFileName,
	}
	format.User = formatCampaignUser

	//* get campaign images for define a list images of single campaign
	images := []GetSpecifiedImageCampaignFormatterResponse{}
	for _, image := range campaign.CampaignImages {
		campaignImage := GetSpecifiedImageCampaignFormatterResponse{}
		campaignImage.ImageUrl = "/" + image.FileName

		isPrimary := false 

		if (image.IsPrimary) {
			isPrimary = true
		}

		campaignImage.IsPrimary = isPrimary

		images = append(images, campaignImage)
	}

	//* assign images with value of array/slice of images
	format.Images = images

	return format
}
