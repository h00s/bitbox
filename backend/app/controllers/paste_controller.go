package controllers

import (
	"errors"
	"net/http"

	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/services"
	"gorm.io/gorm"
)

type PastesController struct {
	raptor.Controller
}

func (pc *PastesController) Pastes() *services.PastesService {
	return pc.Services["PastesService"].(*services.PastesService)
}

func (hc *PastesController) Get(c *raptor.Context) error {
	shortID := c.Params("id")
	paste, err := hc.Pastes().Get(shortID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON("Not found", http.StatusNotFound)
		} else {
			return c.JSON("Internal error", http.StatusInternalServerError)
		}
	}
	return c.JSON(paste)
}
