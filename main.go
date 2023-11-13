package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Joueur struct {
	pseudo  string
	score   int               //Score du joueur
	niv     string            //Choix du niveau (1= niveau 1 etc... jusqu'à 12)
	word    Mot               // Le mot que le mec a
	test    string            //La lettre qu'il veut testé
	win     bool              //S'il a gagné ou pas
	lst     []string          //La liste de lettre qu'il a utilisé
	ind     map[string]string //clé est une lettre et la valeur un message d'indice sur cette lettre
	check   bool              //check si on a réussi a guess une lettre dans le mot (pour l'html pratique)
	message string            //Message affiché selon les cas
}

type Mot struct {
	answer string //Le mot qu'il doit deviner
	gs     string //Le mot qu'il devine (en underscore)
}

var player Joueur = Joueur{} //Déclaration global du joueur

func main() {
	temp, err := template.ParseGlob("./temp/*.html") //Prend tous les .html du dossier template
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur %s", err.Error()))
	}

	http.HandleFunc("/identification", func(w http.ResponseWriter, r *http.Request) { //Pour la route identification
		temp.ExecuteTemplate(w, "identification", nil)
	})

	http.HandleFunc("/treatment/identification", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre cette fonction sert à récupérer les données envoyées par l'utilisateur
		player.pseudo = r.FormValue("pseudo")
	})

	http.HandleFunc("/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
		temp.ExecuteTemplate(w, "niveau", nil)
	})
	http.HandleFunc("/treatment/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.niv = r.FormValue("niveau")
		player.Count()

	})

	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) { //Pour la route jeu
		if len(player.test) > 1 {
			player.TestWord()
		} else if IsInWord(player.word.answer, player.test) {
			if IsInWord(player.word.gs, player.test) {
				player.message = "Vous avez déjà essayez cette lettre"
			} else {
				player.check = true
				player.GuessLetter()
			}
			if player.IsUnderscore() {
				player.win = true
			}
		} else {
			if IsInList(player.lst, player.test) {
				player.message = "Vous avez déjà essayez cette lettre"
			} else {
				player.lst = append(player.lst, player.test)
				player.score--
				player.message = "Mauvaise lettre"
			}
		}
		temp.ExecuteTemplate(w, "jeu", player)
	})

	http.HandleFunc("/treatment/jeu", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.check = false
		if player.win {
			http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)
		} else {
			http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
		}
	})

	http.HandleFunc("/resultat", func(w http.ResponseWriter, r *http.Request) { //Pour la route résultat
		temp.ExecuteTemplate(w, "resultat", nil)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
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

func (p *Joueur) Count() string { //Va mettre le mot que le mec doit deviner avec des underscores
	for n := 0; n < len(p.word.answer); n++ {
		p.word.gs += "_ "
	}
	return p.word.gs
}

func (p *Joueur) indice() {
	p.GuessLetter()
	if p.IsUnderscore() {
		p.win = true
	}
}

func (p *Joueur) letterAleatory() string { //Donne une lettre aléatoire de la réponse
	var w string
	var ale int
	ale = rand.Intn(len(p.word.answer))
	w = string(p.word.answer[ale])
	if IsInList(p.lst, w) {
		w = p.letterAleatory()
	}
	return w
}

func (p *Joueur) GuessLetter() { //Met la lettre que le mec a deviné dans le mot underscore
	p.lst = append(p.lst, p.test)
	for i, t := range p.word.answer {
		if string(t) == p.test {
			p.check = true
			slc := TransformString(p.word.gs)
			slc[i*2] = p.test
			p.word.gs = TransformSlice(slc)
		}
	}
}

func (p *Joueur) TestWord() { //Test si le mot que le mec a rentré est la réponse
	if p.word.answer == p.test {
		player.win = true
	} else {
		p.score -= 2
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
	for _, c := range p.word.gs {
		if string(c) == "_" {
			return false
		}
	}
	return true
}
