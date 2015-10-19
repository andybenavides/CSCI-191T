package main

import (
	"net/http"
	"github.com/golang/appengine"
	"github.com/golang/appengine/datastore"
	"fmt"
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
	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(res, `
	<h2>Hello there...</h2>
		<form method="POST" action="/new_person/">
			Enter your name<input type="text" name="name">
			Enter your birth year<input type="text" name="birthYear">
			<input type="submit">
		</form>
		<form method="POST" action="/view/">
			Click to view birthyears <input type="submit">
		</form>
		`)
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