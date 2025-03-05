package handlers

import (
	"log"
	"net/http"

	"github.com/Githaiga22/forum/backend/repositories"
	"github.com/Githaiga22/forum/backend/util"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("method not allowed")
		util.ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	cookie, err := getSessionID(r)
	if err != nil {
		log.Println("Invalid Session:", err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = repositories.DeleteSession(cookie)
	if err != nil {
		log.Println("error deleting session:", err)
		util.ErrorHandler(w, "An Unexpected Error Occurred. Try Again Later", http.StatusInternalServerError)
		return
	}
	delete(SessionStore, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
