package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/models"
)

type PastesService struct {
	raptor.Service
}

func (ps *PastesService) Get(id uint) (models.PublicPaste, error) {
	var paste models.Paste
	if err := ps.DB.First(&paste, id).Error; err != nil {
		return paste.ToPublicPaste(), err
	}
	return paste.ToPublicPaste(), nil
}

func (ps *PastesService) GetByShortID(shortID string) (models.PublicPaste, error) {
	return ps.Get(models.IDFromShortURI(shortID))
}

func (ps *PastesService) Create(paste models.Paste) (models.PublicPaste, error) {
	err := ps.DB.
		Select(models.PastePermittedParams).
		Create(&paste).Error
	if err != nil {
		return paste.ToPublicPaste(), err
	}
	return ps.Get(paste.ID)
}

func (ps *PastesService) Update(shortID string, paste models.Paste) (models.PublicPaste, error) {
	id := models.IDFromShortURI(shortID)
	err := ps.DB.
		Select(models.PastePermittedParams).
		Model(&models.Paste{}).
		Where("id = ?", id).
		Updates(paste).Error
	if err != nil {
		return paste.ToPublicPaste(), err
	}
	return ps.Get(id)
}
