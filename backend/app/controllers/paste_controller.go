package controllers

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/app/models"
)

type PastesController struct {
	raptor.Controller
}

func (hc *PastesController) Get(c *raptor.Context) error {
	return c.JSON(models.Paste{})
}
