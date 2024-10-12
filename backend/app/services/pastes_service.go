package services

import (
	"context"
	"database/sql"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/bitbox/app/models"
	"github.com/uptrace/bun"
)

type PastesService struct {
	raptor.Service
}

func (ps *PastesService) Get(id int64) (models.PublicPaste, error) {
	var paste models.Paste
	err := ps.DB.Conn().(*bun.DB).
		NewSelect().
		Model(&paste).
		Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			return paste.ToPublicPaste(), raptor.NewErrorNotFound("Paste not found")
		}
		return paste.ToPublicPaste(), raptor.NewErrorInternal(err.Error())
	}
	return paste.ToPublicPaste(), nil
}

func (ps *PastesService) GetByShortID(shortID string) (models.PublicPaste, error) {
	return ps.Get(models.IDFromShortURI(shortID))
}

func (ps *PastesService) Create(paste models.Paste) (models.PublicPaste, error) {
	_, err := ps.DB.Conn().(*bun.DB).
		NewInsert().
		Model(&paste).
		Column(models.PastePermittedParams...).
		Returning("id").
		Exec(context.Background())
	if err != nil {
		return paste.ToPublicPaste(), raptor.NewErrorInternal(err.Error())
	}
	return ps.Get(paste.ID)
}

func (ps *PastesService) Update(shortID string, paste models.Paste) (models.PublicPaste, error) {
	id := models.IDFromShortURI(shortID)
	_, err := ps.DB.Conn().(*bun.DB).
		NewUpdate().
		Model(&paste).
		Column(models.PastePermittedParams...).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return paste.ToPublicPaste(), raptor.NewErrorInternal(err.Error())
	}
	return ps.Get(id)
}
