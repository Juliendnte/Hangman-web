package main

import (
	routeur "hangman/routeur"
	initTemp "hangman/temp"
)

func main() {
	initTemp.InitTemplate() //Init des templates
	routeur.InitServe()     //Init des routes
}
