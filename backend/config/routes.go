package config

import "github.com/go-raptor/raptor"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1/",
			raptor.Route("GET", "pastes/:id", "PastesController", "Get"),
			raptor.Route("POST", "pastes", "PastesController", "Create"),
		),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
