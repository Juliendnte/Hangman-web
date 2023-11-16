package controller

import (
	"bufio"
	initTemp "hangman/temp"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func Niveau(w http.ResponseWriter, r *http.Request) { //Pour la route niveau
	initTemp.Temp.ExecuteTemplate(w, "niveau", nil)
}

func InitNiv(w http.ResponseWriter, r *http.Request) { //Pour le traitement d'une route a une autre
	if r.Method != http.MethodPost {
      	http.Redirect(w, r, "/identification", http.StatusMovedPermanently)
    }
	player.Init()
	player.Niv = r.FormValue("Niveau")
	player.Word.Answer = ToLower(WriteWord("mot/" + player.Niv + ".txt"))
	player.Count()
	http.Redirect(w, r, "/jeu", http.StatusMovedPermanently)
}

func (p *Joueur) Init() {
	p.Word.Gs = ""
	p.Win = false
	p.Test = ""
	p.Hangman.Check = false
	p.Hangman.Message = ""
	p.Word.Answer = ""
	p.Hangman.Score = 0
	p.Lst = nil
	p.Hangman.Img = "p0.png"
	str :="Petit indice: Une des lettres que tu cherche"
    p.Hangman.Ind = map[string]string{"a": str+" est la lettre de l'étonnement",
        "b": str+" est la première lettre du fabriquant de pain",
        "c": str+" est la première lettre du surnom d'Alex",
        "d": str+" est l'une de nos initiale",
        "e": str+" est la lettre la plus utilisé",
        "f": str+" est la lettre la plus basse à une évaluation d'un anglophone",
        "g": str+" est la première lettre de l'animal le plus rapide",
        "h": str+" est une lettre invisible",
        "i": str+" est un indice",
        "j": str+" commence chaque jour",
        "k": str+" est la lettre de l'animal de l'australie",
        "m": str+" est la lettre de ta génitrice",
        "n": str+" est la lettre du refus",
        "o": str+" est la lettre de la comtemplation",
        "p": str+" est la lettre de ton géniteur",
        "q": str+" est la lettre des fesses",
        "r": str+" est la lettre qui 21% dioxygène 78% diazote et 1% de gazs rares",
        "s": str+" c'est le j ",
        "t": str+" ",
        "u": str+" is you",
        "v": str+" est le signe du dems",
        "w": str+" est la lettre la moins utilisé",
        "x": str+" est une lettre de coquin",
        "y": str+" cabre",
        "z": str+" est en bas"}
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
	for n := 0; n < len(p.Word.Answer); n++ {
	    if string(p.Word.Answer[n])=="-"{
	        p.Word.Gs+="- "
	    }else{
	        p.Word.Gs += "_ "
	    }

	}
	return p.Word.Gs
}