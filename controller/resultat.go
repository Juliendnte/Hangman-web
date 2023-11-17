package controller

import (
    initTemp "hangman/temp"
    hang "hangman/Hangman"
    "net/http"
)

func Resultat(w http.ResponseWriter, r *http.Request) { //Pour la route résultat
	initTemp.Temp.ExecuteTemplate(w, "resultat", player)
}
