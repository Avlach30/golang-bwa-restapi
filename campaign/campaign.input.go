package campaign

type GetSpecifiedCampaignInput struct {
	ID int `uri:"id" binding:"required"` //* uri key for get value of request params
}