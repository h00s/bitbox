package components

import (
	"github.com/go-raptor/raptor/v4"
)

func New() *raptor.Components {
	return &raptor.Components{
		Services:    Services(),
		Middlewares: Middlewares(),
		Controllers: Controllers(),
	}
}
