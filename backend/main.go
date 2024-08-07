package main

import (
	"github.com/go-raptor/raptor/v2"
	"github.com/h00s/bitbox/config/initializers"
)

func main() {
	r := raptor.NewRaptor()

	r.Init(initializers.App(r.Utils.Config))

	r.Listen()
}
