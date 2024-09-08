package migrate

import (
	"context"

	"github.com/go-raptor/raptor/v3"
	"github.com/h00s/bitbox/app/models"
)

func AddPaste(db *raptor.DB) error {
	_, err := db.NewCreateTable().
		Model(&models.Paste{}).
		IfNotExists().
		Exec(context.Background())

	return err
}
