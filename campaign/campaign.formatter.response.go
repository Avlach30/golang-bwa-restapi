package campaign

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
