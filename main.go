package main

import (
	//h "hangman/go"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Joueur struct {
	pseudo string
	score  int      //Score du joueur
	niv    string      //Choix du niveau (1= niveau 1 etc... jusqu'à 12)
	word   Mot      // Le mot que le mec a
	test   string   //La lettre qu'il veut testé
	win    bool     //S'il a gagné ou pas
	lst    []string //La liste de lettre qu'il a utilisé
	ind    int      //Pour savoir s'il peut encore utilisé des indices
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
		temp.ExecuteTemplate(w, "identification", player)
	})

	http.HandleFunc("/treatment/identification", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre cette fonction sert à récupérer les données envoyées par l'utilisateur
		player.pseudo = r.FormValue("pseudo")
	})

	http.HandleFunc("/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
		temp.ExecuteTemplate(w, "niveau", player)
	})
	http.HandleFunc("/treatment/niveau", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.niv = r.FormValue("niveau")
	})

	http.HandleFunc("/jeu", func(w http.ResponseWriter, r *http.Request) { //Pour la route jeu
		temp.ExecuteTemplate(w, "jeu", player)
	})

	http.HandleFunc("/treatment/jeu", func(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
		player.test = r.FormValue("lettre")
	})

	http.HandleFunc("/resultat", func(w http.ResponseWriter, r *http.Request) { //Pour la route résultat
		temp.ExecuteTemplate(w, "resultat", player)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)
}
