package Mail

import (
	"net/http"
	"html/template"
	"github.com/golang/appengine"
	"log"
	"github.com/golang/appengine/datastore"
	"github.com/golang/appengine/mail"
	"fmt"
)

type Email struct{
	Name 		string
	EmailAddr 	string
}

func init(){
	http.HandleFunc("/", mainFunc)
	http.HandleFunc("/send/", sendFunc)
	http.HandleFunc("/sent/", sentFunc)
}

func mainFunc(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		addr := req.FormValue("email")
		name := req.FormValue("name")

		c := appengine.NewContext(req)
		key := datastore.NewKey(c, "Email", addr, 0, nil)

		person := &Email{
			Name: name,
			EmailAddr: addr,
		}

		_, err := datastore.Put(c, key, person)
		if err != nil {
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

func sendFunc(res http.ResponseWriter, req *http.Request){
	var emails []string
	if req.Method == "POST" {
		ctx := appengine.NewContext(req)
		q := datastore.NewQuery("Email").Order("Name")

		iterator := q.Run(ctx)
		for {
			var person Email
			_, err := iterator.Next(&person)
			if err == datastore.Done {
				break
			}else if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			emails = append(emails, person.EmailAddr)
		}
	}

	tpl, err := template.ParseFiles("saved.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "saved.html", emails)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

func sentFunc(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		ctx := appengine.NewContext(req)

		q := datastore.NewQuery("Email").Order("Name")

		var emails []Email
		_, err := q.GetAll(ctx, &emails)
		if err != nil{
			err.Error()
		}


		tplt, err := template.ParseFiles("saved.html")
		if err != nil {
			log.Fatalln(err)
		}

		for _, email := range emails {
			name := email.Name
			email := email.EmailAddr

			msg := &mail.Message{
				Sender: 	"neworganization.net Support <andybenavides.22@gmail.com>",
				To:    		[]string{email},
				Subject:	"Welcome to the club " + name + "!",
				HTMLBody:   fmt.Sprint(tplt),
			}
			if err := mail.Send(ctx, msg); err != nil {
				ctx.Err()
			}
		}
	}
	tpl, err := template.ParseFiles("sent.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "sent.html", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}
}

