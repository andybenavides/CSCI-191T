package main

import (
	"html/template"
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/submitPage.html", submitted)
	http.ListenAndServe(":9000", nil)
}

func mainPage(res http.ResponseWriter, req *http.Request){
	fn := "firstName"
	ln := "lastName"
	firstName := req.FormValue(fn)
	lastName := req.FormValue(ln)
	fullName := firstName + " " + lastName

	tpl, err := template.ParseFiles("main.html")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(res, fullName)
	if err != nil{
		log.Fatalln(err)
	}
}

func submitted(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("submitPage.html")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.Execute(res, nil)
}