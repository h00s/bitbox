package config

import "github.com/go-raptor/raptor/v3"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("api/v1",
			raptor.Get("pastes/:id", "Pastes#Get"),
			raptor.Patch("pastes/:id", "Pastes#Update"),
			raptor.Post("pastes", "Pastes#Create"),
		),
	)
}
