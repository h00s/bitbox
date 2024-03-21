package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/models"
)

type PastesService struct {
	raptor.Service
}

func (p *PastesService) Get(id uint) (models.PublicPaste, error) {
	var paste models.Paste
	if err := p.DB.First(&paste, id).Error; err != nil {
		return paste.ToPublicPaste(), err
	}
	return paste.ToPublicPaste(), nil
}

func (p *PastesService) GetByShortID(shortID string) (models.PublicPaste, error) {
	return p.Get(models.IDFromShortURI(shortID))
}

func (p *PastesService) Create(paste models.Paste) (models.PublicPaste, error) {
	if err := p.DB.Select(models.PastePermittedParams).Create(&paste).Error; err != nil {
		return paste.ToPublicPaste(), err
	}
	return p.Get(paste.ID)
}

func (p *PastesService) Update(shortID string, paste models.Paste) (models.PublicPaste, error) {
	id := models.IDFromShortURI(shortID)
	if err := p.DB.Select(models.PastePermittedParams).Model(&models.Paste{}).Where("id = ?", id).Updates(paste).Error; err != nil {
		return paste.ToPublicPaste(), err
	}
	return p.Get(id)
}
