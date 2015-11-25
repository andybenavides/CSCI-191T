package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Me struct{
	Name string
	Age int
}

func CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	andy := Me{
		Name: "Andy",
		Age: 24,
	}
	tpl.ExecuteTemplate(res, "homePage.html", andy)
}
