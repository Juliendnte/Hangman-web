package hangman

import (
	"math/rand"
	"strconv"
)

type Joueur struct {
	Pseudo   string
	Mdp      string   //Mot de passe du joueur
	ScoreG   int      //Score du joueur
	Niv      string   //Choix du niveau (1= niveau 1 etc... jusqu'à 12)
	NivScore int      //je recup le niveau et je le transforme en int pour faire le score avec
	Word     Mot      // Le mot que le mec a
	Win      bool     //Pour savoir s'il a win
	Test     string   //La lettre qu'il veut testé
	Lst      []string //La liste de lettre qu'il a utilisé
	Hangman  Site
}

type Site struct {
	Ind     map[string]string //clé est une lettre et la valeur un message d'indice sur cette lettre
	Score   int               //Score du hangman /6
	Check   bool              //check si on a réussi a guess une lettre dans le mot (pour l'html pratique)
	Message string            //Message affiché selon les cas
	Img     string            //Url pour l'image
}

type Mot struct {
	Answer string //Le mot qu'il doit deviner
	Gs     string //Le mot qu'il devine (en underscore)
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
		"t": "Une des lettres que tu cherche est la lettre",
		"u": "Une des lettres que tu cherche est la lettre",
		"v": "Une des lettres que tu cherche est la lettre",
		"w": "Une des lettres que tu cherche est la lettre",
		"x": "Une des lettres que tu cherche est la lettre",
		"y": "Une des lettres que tu cherche est la lettre",
		"z": "Une des lettres que tu cherche est la lettre"}
}

func (p *Joueur) ImgHangman() {
	if p.Hangman.Score > 6 {
		p.Hangman.Score = 6
	}
	p.Hangman.Img = "p" + strconv.Itoa(p.Hangman.Score) + ".png"
}

// func (p *Joueur) indice() {
// 	p.GuessLetter()
// 	if p.IsUnderscore() {
//		player.Win = true
// 		http.Redirect(w, r, "/resultat", http.StatusMovedPermanently)

// 	}
// }

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

func (p *Joueur) Count() string { //Va mettre le mot que le mec doit deviner avec des underscores
	for n := 0; n < len(p.Word.Answer); n++ {
		p.Word.Gs += "_ "
	}
	return p.Word.Gs
}

func (p Joueur) IsUnderscore() bool { //On regarde s'il y a encore des underscores dans le mot
	for _, c := range p.Word.Gs {
		if string(c) == "_" {
			return false
		}
	}
	return true
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
func Append(lst []string, s string) []string { //Append sans occurence dans la liste
	if !(IsInList(lst, s)) {
		lst = append(lst, s)
	}
	return lst
}
