package controller

import (
	hang "hangman/Hangman"
	initTemp "hangman/temp"
	"net/http"
)

func Niveau(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
	if hang.Player.Hangman.Url != "/niveau" {
		http.Redirect(w, r, hang.Player.Hangman.Url, http.StatusMovedPermanently)
	}
	initTemp.Temp.ExecuteTemplate(w, "niveau", nil)
}

func InitNiv(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
	}
	hang.Player.Init() //Ré-initialise mes valeurs quand je reviens sur la page niveau
	hang.Player.Niv = r.FormValue("Niveau")
	hang.Player.Word.Answer = hang.ToLower(hang.WriteWord("mot/" + hang.Player.Niv + ".txt")) //Prend un mot aléatoire du mode choisi et le minusculise
	hang.Player.Count()                                                                       //Init le mot underscore
	hang.Player.Hangman.Url = "/jeu"
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}
