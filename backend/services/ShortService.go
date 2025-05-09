package services

import (
	"github.com/google/uuid"
	"pendek.in/config"
	"pendek.in/models"
	"pendek.in/utils"
)


func GetShortById(id string) (*models.ShortModel, error) {
	var short models.ShortModel
	result := config.DbClient.First(&short, "short_url = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &short, nil
}


func CreateShortService(data models.ShortModelPayload) (*models.ShortModel, error) {
	var resultCreate models.ShortModel
	shortResult := utils.GeneratorURL(data.Url)

	resultCreate = models.ShortModel{
		Id: uuid.New().String(),
		OriginUrl: data.Url,
		ShortUrl: shortResult,
	}

	result := config.DbClient.Create(&resultCreate)
	if result.Error != nil {
		return nil, result.Error
	}
	
	return &resultCreate,nil
}