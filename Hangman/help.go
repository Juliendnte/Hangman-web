package Hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
)

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
	if p.Hangman.Score > 11 {
		return
	}
	p.Hangman.Img = "tarot" + strconv.Itoa(p.Hangman.Score) + ".jpg"
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

func (p *Joueur) Init() {
	p.Word.Gs = ""
	p.Win = false
	p.Test = ""
	p.Hangman.Check = false
	p.Hangman.Message = ""
	p.Word.Answer = ""
	p.Hangman.Score = 0
	p.Lst = nil
	p.Hangman.Img = "tarot0.jpg"
	str := "Petit indice: Une des lettres que tu cherche"
	p.Hangman.Ind = map[string]string{"a": str + " est la lettre de l'étonnement",
		"b": str + " est la première lettre du fabriquant de pain",
		"c": str + " est la première lettre du surnom d'Alex",
		"d": str + " est l'une de nos initiale",
		"e": str + " est la lettre la plus utilisé",
		"f": str + " est la lettre la plus basse à une évaluation d'un anglophone",
		"g": str + " est la première lettre de l'animal le plus rapide",
		"h": str + " est une lettre invisible",
		"i": str + " est un indice",
		"j": str + " commence chaque jour",
		"k": str + " est la lettre de l'animal de l'australie",
		"m": str + " est la lettre de ta génitrice",
		"n": str + " est la lettre du refus",
		"o": str + " est la lettre de la comtemplation",
		"p": str + " est la lettre de ton géniteur",
		"q": str + " est la lettre des fesses",
		"r": str + " est la lettre qui 21% dioxygène 78% diazote et 1% de gazs rares",
		"s": str + " c'est le j ",
		"t": str + " ",
		"u": str + " is you",
		"v": str + " est le signe du dems",
		"w": str + " est la lettre la moins utilisé",
		"x": str + " est une lettre de coquin",
		"y": str + " cabre",
		"z": str + " est en bas"}
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
		if string(p.Word.Answer[n]) == "-" {
			p.Word.Gs += "- "
		} else {
			p.Word.Gs += "_ "
		}

	}
	return p.Word.Gs
}
