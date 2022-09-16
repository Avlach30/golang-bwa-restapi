package campaign

import (
	"campaigns-restapi/auth"
	"time"
)

type Campaign struct {
	ID               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage //* Implement relations with CampaignImages return an array
	User             auth.User //* Implement relations with User
}

type CampaignImage struct {
	ID         int
	CampaignId int
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
