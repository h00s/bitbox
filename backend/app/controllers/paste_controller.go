package controllers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/services"
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
		return c.JSON(err)
	}
	return c.JSON(paste)
}
