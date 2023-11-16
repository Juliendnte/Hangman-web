package controller

import (
	initTemp "hangman/temp"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
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
	}else if player.Hangman.Score > 6 {
		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
	}
	if player.Hangman.Score==6 && IndCheck==false{
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



func NivToScore(n string) {
	switch n {
	case "mot3lettres":
		player.NivScore = 1
	case "mot4lettres":
		player.NivScore = 2
	case "mot5lettres":
		player.NivScore = 3
	case "mot6lettres":
		player.NivScore = 4
	case "mot7lettres":
		player.NivScore = 5
	case "mot8lettres":
		player.NivScore = 6
	case "mot9lettres":
		player.NivScore = 7
	case "mot10lettres":
		player.NivScore = 8
	case "motpenduanglais":
		player.NivScore = 9
	case "mot+10lettres":
		player.NivScore = 10
	case "multilettres":
		player.NivScore = 11
	case "impossible":
		player.NivScore = 12
	default:
		return
	}
}

func (p *Joueur) ImgHangman() {
	if p.Hangman.Score > 6 {
        return
	}
	p.Hangman.Img = "p" + strconv.Itoa(p.Hangman.Score) + ".png"
}


func (p *Joueur) LetterAleatory() string { //Donne une lettre aléatoire de la réponse
	var w string
	var ale int
	ale = rand.Intn(len(p.Word.Answer))
	w = string(p.Word.Answer[ale])
	if IsInList(p.Lst, w) {
		w = p.LetterAleatory()
	}
	return w
}

func (p *Joueur) GuessLetter() { //Met la lettre que le mec a deviné dans le mot underscore
	p.Lst = Append(p.Lst, p.Test)
	for i, t := range p.Word.Answer {
		if string(t) == p.Test {
			p.Hangman.Check = true
			slc := TransformString(p.Word.Gs)
			slc[i*2] = p.Test
			p.Word.Gs = TransformSlice(slc)
		}
	}
}

func (p *Joueur) TestWord() bool { //Test si le mot que le mec a rentré est la réponse
	if p.Word.Answer == p.Test {
		return true
	} else {
		p.Hangman.Message = "Ce n'est pas le bon mot"
		p.Hangman.Score += 2
		return false
	}
}

func (p Joueur) IsUnderscore() bool { //On regarde s'il y a encore des underscores dans le mot
	for _, c := range p.Word.Gs {
		if string(c) == "_" {
			return false
		}
	}
	return true
}


func IsInList(lst []string, s string) bool { // on regarde si une lettre est dans la liste ou pas
	for _, c := range lst {
		if string(c) == s {
			return true
		}
	}
	return false
}

func TransformString(s string) []string { //Met un mot en []string
	slice := []string{}
	for _, c := range s {
		slice = append(slice, string(c))
	}
	return slice
}

func TransformSlice(s []string) string { //Met un []string en mot
	var str string
	for _, c := range s {
		str += c
	}
	return str
}
func Append(lst []string, s string) []string { //Append sans occurence dans la liste
	if !(IsInList(lst, s)) {
		lst = append(lst, s)
	}
	return lst
}


func IsInWord(word, s string) bool { // on regarde si une lettre est dans le mot ou pas
	for _, l := range word {
		if string(l) == s {
			return true // si ça y est tu peux te le mettre dans le trou (com réalisé par Nath)
		}
	}
	return false
}
