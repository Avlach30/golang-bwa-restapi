package campaign

import "campaigns-restapi/auth"

type GetSpecifiedCampaignInput struct {
	ID int `uri:"id" binding:"required"` //* uri key for get value of request params
}

type CreateNewCampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description"  binding:"required"`
	Description      string `json:"description"  binding:"required"`
	GoalAmount       int    `json:"goal_amount"  binding:"required"`
	Perks            string `json:"perks"  binding:"required"`
	User             auth.User
}
