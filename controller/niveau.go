package controller

import (
	initTemp "hangman/temp"
	hang "hangman/Hangman"
	"net/http"
)

func Niveau(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
	if hang.Player.Hangman.Url !="/niveau"{
		http.Redirect(w, r, hang.Player.Hangman.Url, http.StatusMovedPermanently)
	}
	initTemp.Temp.ExecuteTemplate(w, "niveau", nil)
}

func InitNiv(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
	if r.Method != http.MethodPost {
      	http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
    }
	hang.Player.Hangman.Url="/jeu"
	hang.Player.Init()
	hang.Player.Niv = r.FormValue("Niveau")
	hang.Player.Word.Answer = hang.ToLower(hang.WriteWord("mot/" + hang.Player.Niv + ".txt"))
	hang.Player.Count()
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}
