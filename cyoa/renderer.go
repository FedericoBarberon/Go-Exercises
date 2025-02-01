package cyoa

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	bookTemplates embed.FS
)

type BookRenderer struct {
	tmpl *template.Template
}

func NewBookRenderer() (*BookRenderer, error) {
	tmpl, err := template.ParseFS(bookTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &BookRenderer{tmpl}, nil
}

func (r *BookRenderer) RenderArc(w io.Writer, arc Arc) error {
	return r.tmpl.ExecuteTemplate(w, "arc.gohtml", arc)
}

func (r *BookRenderer) Render404(w io.Writer, title string) error {
	return r.tmpl.ExecuteTemplate(w, "404.gohtml", title)
}
