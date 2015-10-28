package main

import (
	"html/template"
	"net/http"
	"log"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"os/user"
	"google.golang.org/appengine/user"
	"github.com/golang/appengine/user"
)

type Person struct{
	first string
	last string
	birth string
}

func init(){
	http.HandleFunc("/", mainPage)
}

func mainPage(res http.ResponseWriter, req *http.Request){
	c := appengine.NewContext(req)
	u, err := user.Current(c)


	if req.Method == "POST" {
		firstName := req.FormValue("firstName")
		lastName := req.FormValue("lastName")
		birthYear := req.FormValue("birthYear")

		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Visitor", firstName, 0, nil)

		person := &Person{
			first: firstName,
			last: lastName,
			birth: birthYear,
		}

		_, err := datastore.Put(ctx, key, person)
		if err != nil{
			http.Error(res, err.Error(), 500)
			return
		}

	}
	tpl, err := template.ParseFiles("main.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "main.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}