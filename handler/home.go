package handler

import (
	"ai-saas/view/home"
	"net/http"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	//return fmt.Errorf("Failed to generate")
	return home.Index().Render(r.Context(), w)
}
