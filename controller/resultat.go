package controller

import (
	hang "hangman/Hangman"
	initTemp "hangman/temp"
	"net/http"
)

func Resultat(w http.ResponseWriter, r *http.Request) { //Pour la route résultat
	if hang.Player.Hangman.Url != "/resultat" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	initTemp.Temp.ExecuteTemplate(w, "resultat", hang.Player)
	hang.Player.Hangman.Url = "/niveau"
}
