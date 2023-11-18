package initTemplate

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template//Variable global qu'on pourra appelé en important le package

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./temp/*.html")//Tous les .html du dossier temp
	if errTemp != nil {
		fmt.Printf("ErrorTemplates : %v", errTemp.Error())
		os.Exit(1)
	}
	Temp = temp
}
