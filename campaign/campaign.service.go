package campaign

import "errors"

type Service interface {
	FindAllCampaigns(userId int) ([]Campaign, error)
	FindSpecifiedCampaign(input GetSpecifiedCampaignInput) (Campaign, error)
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
