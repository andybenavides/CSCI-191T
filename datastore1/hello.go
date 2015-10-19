package main

import (
	"net/http"
	"github.com/golang/appengine"
	"github.com/golang/appengine/datastore"
	"fmt"
	"html/template"
	"log"
)

func init(){
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/view/", view)
}

type Person struct{
	Name string
	BirthYear string
}

func handleIndex(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST"{
		name := req.FormValue("name")
		year := req.FormValue("birthYear")

		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Person", name, 0, nil)

		person := &Person{
			Name: name,
			BirthYear: year,
		}

		_, err := datastore.Put(ctx, key, person)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}
	tpl, err := template.ParseFiles("home.html")
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "home.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func view(res http.ResponseWriter, req *http.Request){
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Person").Order("BirthYear")

	html := ""

	iterator := q.Run(ctx)
	for{
		var person Person
		_, err := iterator.Next(&person)
		if err == datastore.Done{
			break
		}else if err != nil{
			http.Error(res, err.Error(), 500)
			return
		}
		html += `
			<dt>` + person.Name + `</dt>
			<dd>` + person.BirthYear + `</dd>`
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `
		<dl>
			`+html+`
		</dl>
		<form action="/">
			Go back <input type="submit">
		</form>`)
}