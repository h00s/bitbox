package models

import (
	"strings"
	"time"
)

const validChars = "bcdfghmnprstvz23456789"

var PastePermittedParams = []string{"title", "content"}

type (
	Pastes       []Paste
	PublicPastes []PublicPaste
)

type Paste struct {
	ID        int64     `json:"id" bun:"id,pk,autoincrement"`
	Title     string    `json:"title" bun:"title,notnull"`
	Content   string    `json:"content" bun:"content,notnull"`
	CreatedAt time.Time `json:"created_at" bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,notnull,default:current_timestamp"`
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

func ShortURIfromID(id int64) string {
	uri := ""
	for id > 0 {
		uri = string(validChars[id%int64(len(validChars))]) + uri
		id = id / int64(len(validChars))
	}
	return uri
}

func IDFromShortURI(uri string) int64 {
	id := int64(0)
	for _, c := range uri {
		id = id*int64(len(validChars)) + int64(strings.Index(validChars, string(c)))
	}
	return id
}
