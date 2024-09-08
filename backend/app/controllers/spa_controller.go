package controllers

import "github.com/go-raptor/raptor/v3"

type SPAController struct {
	raptor.Controller
}

func (hc *SPAController) Index(c *raptor.Context) error {
	return c.File("public/index.html")
}
