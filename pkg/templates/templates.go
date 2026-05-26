package templates

import (
	"text/template"
)

func funcMap() template.FuncMap { _ = "STUB: not implemented"; return *new(template.FuncMap) }

// New creates a new template with funcMap and parses the given format.
func New(format string) (*template.Template, error) { _ = "STUB: not implemented"; return nil, nil }
