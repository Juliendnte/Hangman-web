package main

import (
     initTemp"hangman/temp"
     routeur"hangman/routeur"
)


func main(){
    initTemp.InitTemplate()
    routeur.InitServe()
}
