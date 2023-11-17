package controller

import (
	initTemp "hangman/temp"
	"net/http"
	"regexp"
)

var IndCheck bool

func Jeu (w http.ResponseWriter, r *http.Request) { //Pour la route jeu
	if player.Test == "" {
		initTemp.Temp.ExecuteTemplate(w, "jeu", player)
		return
	}
	if len(player.Test) > 1 {
		if player.TestWord() {//return true si le mot qu'il a envoyé est égale au mot
			player.Win = true
			NivToScore(player.Niv)
            player.ScoreG = player.ScoreG + player.NivScore*10
			player.ImgHangman()//Set l'url pour l'affichage du hangman
			http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
		}
	} else if IsInWord(player.Word.Answer, player.Test) {
		if IsInList(player.Lst, player.Test) {//return true si la lettre est dans la liste
			player.Hangman.Check = false
			player.Hangman.Message = "Vous avez déjà essayez cette lettre"
		} else {
			player.Hangman.Message = "Bien trouvé"
			player.Lst = Append(player.Lst, player.Test)
			player.Hangman.Check = true
			player.GuessLetter()//Met la lettre dans le mot avec les underscore
		}
		} else {
		player.Hangman.Check = false
		if IsInList(player.Lst, player.Test) {
			player.Hangman.Message = "Vous avez déjà essayez cette lettre"
		} else {
			player.Lst = Append(player.Lst, player.Test)
			player.Hangman.Score++
			player.Hangman.Message = "Mauvaise lettre"
			player.ImgHangman()
		}
	}
	if player.IsUnderscore() {//Return true s'il n'y a plus d'underscore dans le mot
		player.Win = true
		NivToScore(player.Niv)
		player.ScoreG = player.ScoreG + player.NivScore*10
		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
	}else if player.Hangman.Score > 11 {
		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
	}
	if player.Hangman.Score==11 && IndCheck==false{
	    IndCheck=true
	    player.Hangman.Message=player.Hangman.Ind[player.LetterAleatory()]
	}
	player.ImgHangman()
	initTemp.Temp.ExecuteTemplate(w, "jeu", player)
}

func InitJeu(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
	if r.Method != http.MethodPost {
       	http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
    }
	player.Test = ToLower(r.FormValue("lettre"))
	checkValue, _ := regexp.MatchString("^[a-zA-Z]{1,25}$", player.Test)
	if !checkValue {
		player.Hangman.Message = "Invalide"
		player.Test = ""
	}
	player.Hangman.Check = false
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}

