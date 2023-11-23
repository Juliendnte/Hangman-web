package controller

import (
	"fmt"
	hang "hangman/Hangman"
	initTemp "hangman/temp"
	"net/http"
	"regexp"
)

var IndCheck bool

func Jeu(w http.ResponseWriter, r *http.Request) { //Pour la route jeu
	if hang.Player.Hangman.Url != "/jeu" {
		http.Redirect(w, r, hang.Player.Hangman.Url, http.StatusMovedPermanently)
	}
	if hang.Player.Test == "" {
		initTemp.Temp.ExecuteTemplate(w, "jeu", hang.Player)
		return
	}
	if len(hang.Player.Test) > 1 { //Si le joueur rentre un mot
		if hang.Player.TestWord() { //return true si le mot de l'user est égale à la réponse
			hang.Player.Win = true
			hang.NivToScore(hang.Player.Niv) //Set le nombre de point que l'utilisateur gagne selon le niveau
			hang.Player.ScoreG = hang.Player.ScoreG + hang.Player.NivScore*10
			hang.Player.ImgHangman() //Set l'url pour l'affichage du hangman
			hang.Player.Hangman.Url = "/resultat"
			http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
		}
		//Si le joueur rentre une lettre
	} else if hang.IsInWord(hang.Player.Word.Answer, hang.Player.Test) { //Si la lettre est dans la réponde
		if hang.IsInList(hang.Player.Lst, hang.Player.Test) { //S'il a déjà utilisé la lettre
			hang.Player.Hangman.Check = false
			hang.Player.Hangman.Message = "Vous avez déjà essayez cette lettre"
		} else {
			hang.Player.Hangman.Message = "Bien trouvé"
			hang.Player.Lst = hang.Append(hang.Player.Lst, hang.Player.Test) //Append ne l'ajoute que si la lettre n'est pas déjà dans sa liste
			hang.Player.Hangman.Check = true
			hang.Player.GuessLetter() //Met la lettre dans le mot avec les underscore
		}
	} else {
		hang.Player.Hangman.Check = false
		if hang.IsInList(hang.Player.Lst, hang.Player.Test) {
			hang.Player.Hangman.Message = "Vous avez déjà essayez cette lettre"
		} else {
			hang.Player.Lst = hang.Append(hang.Player.Lst, hang.Player.Test)
			hang.Player.Hangman.Score++
			hang.Player.Hangman.Message = "Mauvaise lettre"
		}
	}
	if hang.Player.IsUnderscore() { //Return true s'il n'y a plus d'underscore dans le mot
		hang.Player.Win = true
		hang.NivToScore(hang.Player.Niv)
		hang.Player.ScoreG = hang.Player.ScoreG + hang.Player.NivScore*10
		hang.Player.Hangman.Url = "/resultat"
		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
	} else if hang.Player.Hangman.Score > 11 { //Cas où il n'a plus de pdv
		hang.Player.Hangman.Url = "/resultat"
		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
	}
	if hang.Player.Hangman.Score == 11 && IndCheck == false { //Pour l'indice
		IndCheck = true
		hang.Player.Hangman.Message = hang.Player.Hangman.Ind[hang.Player.LetterAleatory()]
		fmt.Println(hang.Player.Hangman.Message)
	}
	hang.Player.ImgHangman() //Set le chemin de img que je renvoie à l'html
	initTemp.Temp.ExecuteTemplate(w, "jeu", hang.Player)
}

func InitJeu(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
	}
	hang.Player.Test = hang.ToLower(r.FormValue("lettre"))
	checkValue, _ := regexp.MatchString("^[a-zA-Z]{1,25}$", hang.Player.Test)
	if !checkValue {
		hang.Player.Hangman.Message = "Invalide"
		hang.Player.Test = ""
	}
	hang.Player.Hangman.Check = false
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}
