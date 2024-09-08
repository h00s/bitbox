package controllers

import (
	"net/http"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/bitbox/app/models"
	"github.com/h00s/bitbox/app/services"
)

type PastesController struct {
	raptor.Controller
	Pastes *services.PastesService
}

func (hc *PastesController) Get(c *raptor.Context) error {
	paste, err := hc.Pastes.GetByShortID(c.Param("id"))
	if err != nil {
		return c.JSONError(raptor.NewErrorNotFound("Paste not found"))
	}
	return c.JSON(paste.ToPublicPaste())
}

func (hc *PastesController) Create(c *raptor.Context) error {
	var paste models.Paste
	if err := c.Bind(&paste); err != nil {
		return c.JSON("Invalid input", http.StatusBadRequest)
	}
	p, err := hc.Pastes.Create(paste)
	if err != nil {
		hc.Log.Error(err.Error())
		return c.JSONError(raptor.NewErrorInternal(err.Error()))
	}
	return c.JSON(p, http.StatusCreated)
}

func (hc *PastesController) Update(c *raptor.Context) error {
	shortID := c.Param("id")
	var paste models.Paste
	if err := c.Bind(&paste); err != nil {
		return c.JSONError(raptor.NewErrorBadRequest("Invalid input"))
	}
	p, err := hc.Pastes.Update(shortID, paste)
	if err != nil {
		hc.Log.Error(err.Error())
		return c.JSONError(raptor.NewErrorInternal(err.Error()))
	}
	return c.JSON(p)
}
