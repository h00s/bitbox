package services

import (
	"context"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/bitbox/app/models"
)

type PastesService struct {
	raptor.Service
}

func (ps *PastesService) Get(id int64) (models.PublicPaste, error) {
	var paste models.Paste
	err := ps.DB.
		NewSelect().
		Model(&paste).
		Where("id = ?", id).
		Scan(context.Background())
	return paste.ToPublicPaste(), err
}

func (ps *PastesService) GetByShortID(shortID string) (models.PublicPaste, error) {
	return ps.Get(models.IDFromShortURI(shortID))
}

func (ps *PastesService) Create(paste models.Paste) (models.PublicPaste, error) {
	_, err := ps.DB.
		NewInsert().
		Model(&paste).
		Column(models.PastePermittedParams...).
		Returning("id").
		Exec(context.Background())
	if err != nil {
		return paste.ToPublicPaste(), err
	}
	return ps.Get(paste.ID)
}

func (ps *PastesService) Update(shortID string, paste models.Paste) (models.PublicPaste, error) {
	id := models.IDFromShortURI(shortID)
	_, err := ps.DB.NewUpdate().
		Model(&paste).
		Column(models.PastePermittedParams...).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return paste.ToPublicPaste(), err
	}
	return ps.Get(id)
}
