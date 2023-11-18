package controller

import (
	hang "hangman/Hangman"
	initTemp "hangman/temp"
	"net/http"
)

func Identification(w http.ResponseWriter, r *http.Request) {
	if hang.Player.Hangman.Url != "" { //Securise la route pour ne pas pouvoir y rentrer de force
		http.Redirect(w, r, hang.Player.Hangman.Url, http.StatusMovedPermanently)
	}
	initTemp.Temp.ExecuteTemplate(w, "identification", nil)
}

func InitId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	hang.Player.Pseudo = r.FormValue("pseudo")
	hang.Player.Hangman.Url = "/niveau"
	http.Redirect(w, r, "/niveau", http.StatusMovedPermanently)
}
