package user

import (
	"net/http"

	"github.com/vakishna/go_crash_course/AuthApp/app"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	template.RenderTemplate(w, "user", session.Values["profile"])
}
