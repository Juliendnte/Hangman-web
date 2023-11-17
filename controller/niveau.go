package controller

import (
	initTemp "hangman/temp"
	"net/http"
)

func Niveau(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
	initTemp.Temp.ExecuteTemplate(w, "niveau", nil)
}

func InitNiv(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
	if r.Method != http.MethodPost {
      	http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
    }
	player.Init()
	player.Niv = r.FormValue("Niveau")
	player.Word.Answer = ToLower(WriteWord("mot/" + player.Niv + ".txt"))
	player.Count()
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}
