package Hangman

type Joueur struct {
	Pseudo     string
	ScoreG     int      //Score du joueur
	Niv        string   //Choix du niveau
	Difficulty int      //Difficulté du niveau
	NivScore   int      //je recup le niveau et je le transforme en int pour faire le score avec
	Word       Mot      // Le mot que le mec a
	Win        bool     //Pour savoir s'il a win
	Test       string   //La lettre qu'il veut testé
	Lst        []string //La liste de lettre qu'il a utilisé
	LstHtml    string   //Lst en string
	Hangman    Site
}

type Site struct {
	Ind     map[string]string //clé est une lettre et la valeur un message d'indice sur cette lettre
	Score   int               //Score du hangman /12
	Check   bool              //check si on a réussi a guess une lettre dans le mot (pour l'html pratique)
	Message string            //Message affiché selon les cas
	Url     string            //Url pour savoir où je suis et où je dois aller
	Img     string            //Le chemin de l'image
}

type Mot struct {
	Answer string //Le mot qu'il doit deviner
	Gs     string //Le mot qu'il devine (en underscore)
}

var Player Joueur //Déclaration global du joueur
