package campaign

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/gosimple/slug"
)

type Service interface {
	FindAllCampaigns(userId int) ([]Campaign, error)
	FindSpecifiedCampaign(input GetSpecifiedCampaignInput) (Campaign, error)
	CreateNewCampaign(input CreateNewCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) FindAllCampaigns(userId int) ([]Campaign, error) {
	if (userId != 0) {
		campaigns, err := service.repository.FindAllByUser(userId)
		if err != nil {
			return campaigns, errors.New("failed to get all campaigns by logged user")
		}

		return campaigns, nil
	}

	campaigns, err := service.repository.FindAll()
	if err != nil {
		return campaigns, errors.New("failed to get all campaigns by logged user")
	}

	return campaigns, nil
}

func (service *service) FindSpecifiedCampaign(input GetSpecifiedCampaignInput) (Campaign, error) {
	
	campaign, err := service.repository.FindSpecifiedCampaign(input.ID)
	if (err != nil) {
		return campaign, errors.New("data not found")
	}

	return campaign, nil
}

func (service *service) CreateNewCampaign(input CreateNewCampaignInput) (Campaign, error) {
	campaign := Campaign{
		Name: input.Name,
		ShortDescription: input.ShortDescription,
		Description: input.Description,
		GoalAmount: input.GoalAmount,
		Perks: input.Perks,
	}

	//* assign userid with input.user.Id
	campaign.UserId = input.User.ID

	//* make unique slug and assign it
	slugRandString := fmt.Sprintf("%s - %s", input.Name, strconv.Itoa(input.User.ID))
	campaign.Slug = slug.Make(slugRandString)

	newCampaign, err := service.repository.Save(campaign)
	if (err != nil) {
		return newCampaign, errors.New("failed to save new campaign to database")
	}

	return newCampaign, nil
}
