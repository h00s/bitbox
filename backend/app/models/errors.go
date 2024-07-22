package models

import "github.com/go-raptor/raptor/v2"

func NewError(code int, message string) raptor.Map {
	return raptor.Map{
		"code":    code,
		"message": message,
	}
}
