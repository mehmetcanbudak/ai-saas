package handler

import (
	Settings "ai-saas/view/settings"
	"net/http"
)

func HandleSettingsIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)
	return render(r, w, Settings.Index((user)))
}
