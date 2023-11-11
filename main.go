package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Joueur struct {
	score int
	val   int //val
	word  Mot
	test  string   //test
	win   bool     //win
	lst   []string //lst
	ind   int      //ind
}

type Mot struct {
	answer string //gs
	gs     string //guess
}

var player Joueur = Joueur{}

func main() {
	temp, err := template.ParseGlob("./temp/*.html")
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur %s", err.Error()))
	}
	http.HandleFunc("/identification", func(w http.ResponseWriter, r *http.Request) {
		page := 0
		temp.ExecuteTemplate(w, "identification", page)
	})
	http.HandleFunc("/treatment", func(w http.ResponseWriter, r *http.Request) {
	})
	http.HandleFunc("/niveau", func(w http.ResponseWriter, r *http.Request) {
		page := 0
		temp.ExecuteTemplate(w, "niveau", page)
	})
	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {
		page := 0
		temp.ExecuteTemplate(w, "jeu", page)
	})
	http.HandleFunc("/resultat", func(w http.ResponseWriter, r *http.Request) {
		page := 0
		temp.ExecuteTemplate(w, "resultat", page)
	})
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
