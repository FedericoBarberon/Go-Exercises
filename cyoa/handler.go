package cyoa

import (
	"net/http"
)

type Handler struct {
	book     Book
	renderer *BookRenderer
}

func NewHandler(book Book) (*Handler, error) {
	renderer, err := NewBookRenderer()
	if err != nil {
		return nil, err
	}
	return &Handler{book, renderer}, nil
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")

	arcTitle := r.URL.Path[1:]

	if arcTitle == "" {
		arcTitle = "intro"
	}

	arc, ok := h.book[arcTitle]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		h.renderer.Render404(w, arcTitle)
		return
	}

	h.renderer.RenderArc(w, arc)
}
