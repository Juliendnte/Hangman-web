package routeur

import (
	"fmt"
	h "hangman/controller"
	"net/http"
	"os"
)

func InitServe() {//Appelle la fonction attitré
	http.HandleFunc("/", h.Identification)
	http.HandleFunc("/treatment/identification", h.InitId)
	http.HandleFunc("/niveau", h.Niveau)
	http.HandleFunc("/treatment/niveau", h.InitNiv)
	http.HandleFunc("/jeu", h.Jeu)
	http.HandleFunc("/treatment/jeu", h.InitJeu)
	http.HandleFunc("/resultat", h.Resultat)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("(http://localhost:8081/) - Server started on port:8081")
	http.ListenAndServe("localhost:8081", nil)
	fmt.Println("Server close on port:8081")
}
