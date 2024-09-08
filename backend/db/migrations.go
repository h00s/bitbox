package db

import (
	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/bitbox/db/migrate"
)

func Migrations() raptor.Migrations {
	return raptor.Migrations{
		1: migrate.AddPaste,
	}
}
