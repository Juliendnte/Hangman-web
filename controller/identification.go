package controller

import (
	"net/http"
	hang "hangman/Hangman"
	initTemp "hangman/temp"
)

func Identification (w http.ResponseWriter, r *http.Request) { //Pour la route identification
	if hang.Player
	initTemp.Temp.ExecuteTemplate(w, "identification", nil)
}

func InitId(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
    	http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
    }
	player.Pseudo = r.FormValue("pseudo")
	player.Mdp = r.FormValue("mot")
	http.Redirect(w, r, "/niveau", http.StatusMovedPermanently)
}