package viewmodel

import (
	"html/template"
	"strings"
)

type Post struct {
	Title, Description string
	Tags               []string
	HTMLBody           template.HTML
}

func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
