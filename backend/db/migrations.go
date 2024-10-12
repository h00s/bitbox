package db

import (
	"github.com/go-raptor/connector/bun/postgres"
	"github.com/h00s/bitbox/db/migrate"
)

func Migrations() postgres.Migrations {
	return postgres.Migrations{
		1: migrate.AddPaste,
	}
}
