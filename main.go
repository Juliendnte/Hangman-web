package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Joueur struct {
	score int//Score du joueur
	val   int //Choix du niveau (1= niveau 1 etc... jusqu'à 12)
	word  Mot // Le mot que le mec a
	test  string   //La lettre qu'il veut testé
	win   bool     //S'il a gagné ou pas
	lst   []string //La liste de lettre qu'il a utilisé
	ind   int      //Pour savoir s'il peut encore utilisé des indices
}

type Mot struct {
	answer string //Le mot qu'il doit deviner
	gs     string //Le mot qu'il devine (en underscore)
}

var player Joueur = Joueur{}//Déclaration global du joueur

func main() {
	temp, err := template.ParseGlob("./temp/*.html")//Prend tous les .html du dossier template
	if err != nil {
		fmt.Println(fmt.Sprintf("Erreur %s", err.Error()))
	}


	http.HandleFunc("/identification", func(w http.ResponseWriter, r *http.Request) {//Pour la route identification
		page := 0
		temp.ExecuteTemplate(w, "identification", page)
	})


	http.HandleFunc("/treatment", func(w http.ResponseWriter, r *http.Request) {//Pour le traitement d'une route a une autre
	})


	http.HandleFunc("/niveau", func(w http.ResponseWriter, r *http.Request) {//Pour la route niveau
		page := 0
		temp.ExecuteTemplate(w, "niveau", page)
	})


	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) {//Pour la route jeu
		page := 0
		temp.ExecuteTemplate(w, "jeu", page)
	})


	http.HandleFunc("/resultat", func(w http.ResponseWriter, r *http.Request) {//Pour la route résultat
		page := 0
		temp.ExecuteTemplate(w, "resultat", page)
	})


	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
