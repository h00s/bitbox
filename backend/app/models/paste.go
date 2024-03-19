package models

import (
	"strings"
	"time"
)

const validChars = "bcdfghmnprstvz23456789"

var PasteOmittedParams = []string{"CreatedAt", "UpdatedAt", "ID"}

type (
	Pastes       []Paste
	PublicPastes []PublicPaste
)

type Paste struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PublicPaste struct {
	ID string `json:"id"`
	Paste
}

func (p *Paste) ToPublicPaste() PublicPaste {
	return PublicPaste{
		ID:    ShortURIfromID(p.ID),
		Paste: *p,
	}
}

func ShortURIfromID(id uint) string {
	uri := ""
	for id > 0 {
		uri = string(validChars[id%uint(len(validChars))]) + uri
		id = id / uint(len(validChars))
	}
	return uri
}

func IDFromShortURI(uri string) uint {
	id := uint(0)
	for _, c := range uri {
		id = id*uint(len(validChars)) + uint(strings.Index(validChars, string(c)))
	}
	return id
}
