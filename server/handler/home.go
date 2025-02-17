package handler

import (
	"github.com/siobh9/first-htmx/components/core"
	"github.com/siobh9/first-htmx/components/home"
	"net/http"
)

// Home handles the home page.
func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	h.html(r.Context(), w, http.StatusOK, core.HTML("Example Site", home.Page()))
}
