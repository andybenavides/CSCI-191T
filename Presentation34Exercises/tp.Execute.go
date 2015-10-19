package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	name := "Andy"

	tpl, err := template.ParseFiles("index.html")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, name)
	if err != nil{
		log.Fatalln(err)
	}

}
