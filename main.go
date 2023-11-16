package main

import (
	"bufio"
	"fmt"
	h "hangman/golang"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
)

var player h.Joueur //Déclaration global du joueur

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
		player.Mdp = r.FormValue("mot")
		http.Redirect(w, r, "/niveau", http.StatusMovedPermanently)
	})

	http.HandleFunc("/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
		temp.ExecuteTemplate(w, "niveau", nil)
	})
	http.HandleFunc("/treatment/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.Init()
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
				player.ImgHangman()
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

	fmt.Println("(http://localhost:8081/niveau) - Server started on port:8081")
	http.ListenAndServe("localhost:8081", nil)
	fmt.Println("Server close on port:8081")
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
