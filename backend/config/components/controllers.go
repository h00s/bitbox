package components

import (
	"github.com/go-raptor/raptor/v4"
	"github.com/h00s/bitbox/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.Controllers{
		&controllers.HealthController{},
	}
}
