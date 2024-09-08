package initializers

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/bitbox/app/services"
)

func Services() raptor.Services {
	return raptor.Services{
		&services.PastesService{},
	}
}
