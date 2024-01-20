package main

import (
	"github.com/go-raptor/raptor"
	"github.com/h00s/bitbox/config"
	"github.com/h00s/bitbox/config/initializers"
)

func main() {
	r := raptor.NewRaptor()

	r.Init(initializers.App())
	r.Routes(config.Routes())

	r.Listen()
}
