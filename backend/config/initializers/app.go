package initializers

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/bitbox/config"
)

func App(c *raptor.Config) *raptor.AppInitializer {
	return &raptor.AppInitializer{
		Routes:      config.Routes(),
		Database:    Database(),
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
