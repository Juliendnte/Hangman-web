package main

import (
     initTemp"hangman/temp"
     routeur"hangman/routeur"
)


func main(){
    initTemp.InitTemplate()//Init des templates
    routeur.InitServe()//Init des routes
}
