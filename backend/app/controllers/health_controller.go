// Package controllers provides the controller for handling health check requests.
package controllers

import (
	"github.com/go-raptor/raptor/v4"
)

type HealthController struct {
	raptor.Controller
}

func (c *HealthController) Get(ctx *raptor.Context) error {
	return ctx.Data("OK")
}
