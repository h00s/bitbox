package controllers

import (
	"errors"
	"net/http"

	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/models"
	"github.com/h00s/bitbox/app/services"
	"gorm.io/gorm"
)

type PastesController struct {
	raptor.Controller
	Pastes *services.PastesService
}

func (hc *PastesController) Get(c *raptor.Context) error {
	paste, err := hc.Pastes.GetByShortID(c.Params("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON("Not found", http.StatusNotFound)
		} else {
			return c.JSON("Internal error", http.StatusInternalServerError)
		}
	}
	return c.JSON(paste)
}

func (hc *PastesController) Create(c *raptor.Context) error {
	var paste models.Paste
	if err := c.BodyParser(&paste); err != nil {
		return c.JSON("Invalid input", http.StatusBadRequest)
	}
	p, err := hc.Pastes.Create(paste)
	if err != nil {
		return c.JSON("Internal error", http.StatusInternalServerError)
	}
	return c.JSON(p, http.StatusCreated)
}

func (hc *PastesController) Update(c *raptor.Context) error {
	shortID := c.Params("id")
	var paste models.Paste
	if err := c.BodyParser(&paste); err != nil {
		return c.JSON("Invalid input", http.StatusBadRequest)
	}
	p, err := hc.Pastes.Update(shortID, paste)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON("Not found", http.StatusNotFound)
		} else {
			return c.JSON("Internal error", http.StatusInternalServerError)
		}
	}
	return c.JSON(p)
}
