package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	andy := User{
		FirstName: "Andy",
		LastName: "Benavides",
		Email: "andybenavides.22@gmail.com",
		UserName: "andybenavides.22",
		Age: 24,
	}
	tpl.ExecuteTemplate(res, "homePage.html", andy)
}
