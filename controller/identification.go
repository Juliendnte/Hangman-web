package controller

import (
	hang "hangman/Hangman"
	initTemp "hangman/temp"
	"net/http"
)

func Identification(w http.ResponseWriter, r *http.Request) { //Pour la route identification
	if hang.Player.Hangman.Url != "" {
		http.Redirect(w, r, hang.Player.Hangman.Url, http.StatusMovedPermanently)
	}
	initTemp.Temp.ExecuteTemplate(w, "identification", nil)
}


func InitId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	hang.Player.Hangman.Url = "/niveau"
	hang.Player.Pseudo = r.FormValue("pseudo")
	http.Redirect(w, r, "/niveau", http.StatusMovedPermanently)
}
