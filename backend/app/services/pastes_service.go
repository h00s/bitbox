package services

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/models"
)

type PastesService struct {
	raptor.Service
}

func (p *PastesService) Get(shortID string) (models.PublicPaste, error) {
	id := models.IDFromShortURI(shortID)
	var paste models.Paste
	if err := p.DB.First(&paste, id).Error; err != nil {
		return paste.ToPublicPaste(), err
	}
	return paste.ToPublicPaste(), nil
}

func (p *PastesService) Create(paste models.Paste) (models.PublicPaste, error) {
	if err := p.DB.Omit(models.PasteOmittedParams...).Create(&paste).Error; err != nil {
		return paste.ToPublicPaste(), err
	}
	return paste.ToPublicPaste(), nil
}