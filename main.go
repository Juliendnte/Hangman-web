package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

type Joueur struct {
	Pseudo  string
	ScoreG  int      //Score du joueur
	Niv     string   //Choix du niveau (1= niveau 1 etc... jusqu'à 12)
	Word    Mot      // Le mot que le mec a
	Win     bool     //Pour savoir s'il a win
	Test    string   //La lettre qu'il veut testé
	Lst     []string //La liste de lettre qu'il a utilisé
	Hangman Site
}

type Site struct {
	Ind     map[string]string //clé est une lettre et la valeur un message d'indice sur cette lettre
	Score   int               //Score du hangman
	Check   bool              //check si on a réussi a guess une lettre dans le mot (pour l'html pratique)
	Message string            //Message affiché selon les cas
	Img     string            //Url pour l'image
}

type Mot struct {
	Answer string //Le mot qu'il doit deviner
	Gs     string //Le mot qu'il devine (en underscore)
}

var player Joueur = Joueur{} //Déclaration global du joueur

func main() {
	temp, err := template.ParseGlob("./temp/*.html") //Prend tous les .html du dossier template
	if err != nil {
		fmt.Printf("Erreur %s", err.Error())
	}

	http.HandleFunc("/identification", func(w http.ResponseWriter, r *http.Request) { //Pour la route identification
		temp.ExecuteTemplate(w, "identification", nil)
	})

	http.HandleFunc("/treatment/identification", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre cette fonction sert à récupérer les données envoyées par l'utilisateur
		player.Pseudo = r.FormValue("pseudo")
	})

	http.HandleFunc("/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
		temp.ExecuteTemplate(w, "niveau", nil)
	})
	http.HandleFunc("/treatment/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.init()
		player.Niv = r.FormValue("Niveau")
		player.Word.Answer = ToLower(WriteWord("mot/" + player.Niv + ".txt"))
		player.Count()
		http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
	})

	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) { //Pour la route jeu
		if player.Test == "" {
			temp.ExecuteTemplate(w, "jeu", player)
			return
		}
		if len(player.Test) > 1 {
			if player.TestWord() {
				player.Win = true
				http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
			}
		} else if IsInWord(player.Word.Answer, player.Test) {
			if IsInList(player.Lst, player.Test) {
			    player.Hangman.Check = false
				player.Hangman.Message = "Vous avez déjà essayez cette lettre"
			} else {
				player.Hangman.Message = "Bien trouvé"
				player.Lst = Append(player.Lst, player.Test)
				player.Hangman.Check = true
				player.GuessLetter()
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
		if player.IsUnderscore() {
			player.Win = true
			http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
		} else if player.Hangman.Score >= 6 {
			http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
		}
		player.ImgHangman()
		temp.ExecuteTemplate(w, "jeu", player)
	})

	http.HandleFunc("/treatment/jeu", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.Test = ToLower(r.FormValue("lettre"))
		checkValue, _ := regexp.MatchString("^[a-zA-Z-]$", player.Test)
		if !checkValue {
			player.Hangman.Message = "Invalide"
			player.Test = ""
		}
		player.Hangman.Check = false
		http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
	})

	http.HandleFunc("/resultat", func(w http.ResponseWriter, r *http.Request) { //Pour la route résultat
		temp.ExecuteTemplate(w, "resultat", player)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8081", nil)
}

func (p *Joueur) init() {
	p.Word.Gs = ""
	p.Win = false
	p.Test = ""
	p.Hangman.Check = false
	p.Hangman.Message = ""
	p.Word.Answer = ""
	p.Hangman.Score = 0
	p.Lst = nil
	p.Hangman.Img = "p0.png"
}

func ReadLines(path string) ([]string, error) { //Met un .txt en []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func WriteWord(path string) string { //Prend un mot aléatoirement d'un .txt
	f, err := ReadLines(path)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	ale := rand.Intn(len(f))
	return f[ale]
}

func Append(lst []string, s string) []string { //Append sans occurence dans la liste
	if !(IsInList(lst, s)) {
		lst = append(lst, s)
	}
	return lst
}

func (p *Joueur) ImgHangman() {
    if p.Hangman.Score>6{
        p.Hangman.Score=6
    }
    p.Hangman.Img="p"+strconv.Itoa(p.Hangman.Score)+".png"
}

// func (p *Joueur) indice() {
// 	p.GuessLetter()
// 	if p.IsUnderscore() {
//		player.Win = true
// 		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)

// 	}
// }

func (p *Joueur) letterAleatory() string { //Donne une lettre aléatoire de la réponse
	var w string
	var ale int
	ale = rand.Intn(len(p.Word.Answer))
	w = string(p.Word.Answer[ale])
	if IsInList(p.Lst, w) {
		w = p.letterAleatory()
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

func ToLower(s string) string { //Minuscilise toutes les lettres d'un mot
	var listf string
	for _, c := range s {
		if c > 64 && c < 91 {
			listf = listf + string(c+32)
		} else {
			listf = listf + string(c)
		}
	}
	return listf
}

func (p *Joueur) Count() string { //Va mettre le mot que le mec doit deviner avec des underscores
	for n := 0; n < len(p.Word.Answer); n++ {
		p.Word.Gs += "_ "
	}
	return p.Word.Gs
}

func IsInWord(word, s string) bool { // on regarde si une lettre est dans le mot ou pas
	for _, l := range word {
		if string(l) == s {
			return true // si ça y est tu peux te le mettre dans le trou (com réalisé par Nath)
		}
	}
	return false
}

func IsInList(lst []string, s string) bool { // on regarde si une lettre est dans la liste ou pas
	for _, c := range lst {
		if string(c) == s {
			return true
		}
	}
	return false
}

func (p Joueur) IsUnderscore() bool { //On regarde s'il y a encore des underscores dans le mot
	for _, c := range p.Word.Gs {
		if string(c) == "_" {
			return false
		}
	}
	return true
}
