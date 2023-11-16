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
    p.Hangman.Ind = map[string]string{"a": "Une des lettres que tu cherche est la lettre de l'étonnement",
        "b": "Une des lettres que tu cherche est la première lettre du fabriquant de pain",
        "c": "Une des lettres que tu cherche est la première lettre du surnom d'Alex",
        "d": "Une des lettres que tu cherche est l'une de nos initiale",
        "e": "Une des lettres que tu cherche est la lettre la plus utilisé",
        "f": "Une des lettres que tu cherche est la lettre la plus basse à une évaluation d'un anglophone",
        "g": "Une des lettres que tu cherche est la première lettre de l'animal le plus rapide",
        "h": "Une des lettres que tu cherche est une lettre invisible",
        "i": "Une des lettres que tu cherche est un indice",
        "j": "Une des lettres que tu cherche commence chaque jour",
        "k": "Une des lettres que tu cherche est la lettre de l'animal de l'australie",
        "m": "Une des lettres que tu cherche est la lettre de ta génitrice",
        "n": "Une des lettres que tu cherche est la lettre du refus",
        "o": "Une des lettres que tu cherche est la lettre de la comtemplation",
        "p": "Une des lettres que tu cherche est la lettre de ton géniteur",
        "q": "Une des lettres que tu cherche est la lettre des fesses",
        "r": "Une des lettres que tu cherche est la lettre qui 21% dioxygène 78% diazote et 1% de gazs rares",
        "s": "Une des lettres que tu cherche c'est le j ",
        "t":"Une des lettres que tu cherche est la lettre",
        "u":"Une des lettres que tu cherche est la lettre",
        "v":"Une des lettres que tu cherche est la lettre",
        "w":"Une des lettres que tu cherche est la lettre",
        "x":"Une des lettres que tu cherche est la lettre",
        "y":"Une des lettres que tu cherche est la lettre",
        "z":"Une des lettres que tu cherche est la lettre"}
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
		p.Word.Gs += "_ "
	}
	return p.Word.Gs
}