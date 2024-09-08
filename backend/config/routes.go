package config

import "github.com/go-raptor/raptor/v3"

func Routes() raptor.Routes {
	return raptor.CollectRoutes(
		raptor.Scope("/api/v1/",
			raptor.Route("GET", "pastes/:id", "PastesController", "Get"),
			raptor.Route("PATCH", "pastes/:id", "PastesController", "Update"),
			raptor.Route("POST", "pastes", "PastesController", "Create"),
		),
		raptor.Route("GET", "*", "SPAController", "Index"),
	)
}
