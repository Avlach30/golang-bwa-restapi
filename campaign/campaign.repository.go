package campaign

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindAllByUser(userId int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repository *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := repository.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if (err != nil) {
		return campaigns, errors.New("failed to get all campaigns")
	}

	return campaigns, nil
}

func (repository *repository) FindAllByUser(userId int) ([]Campaign, error) { 
	var campaigns []Campaign

	//* Method preload untuk implementasi relasional. param ke 2 opsional untuk kriteria data
	err := repository.db.Where("user_id = ?", userId).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if (err != nil) {
		return campaigns, errors.New("failed to get all campaigns by logged user")
	}

	return campaigns, nil
	
}