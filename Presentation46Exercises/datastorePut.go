package main

import (
	"net/http"
	"fmt"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type Word struct{
	Term string
	Def string
}

func init(){
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/words/", getWord)
	http.ListenAndServe(":9000", nil)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		term := req.FormValue("term")
		def := req.FormValue("def")

		ctx := appengine.NewContext(req)
		key := datastore.NewKey(ctx, "Word", term, 0, nil)

		entity := &Word{
			Term: term,
			Def: def,
		}

		_, err := datastore.Put(ctx, key, entity)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
	}

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
	<form method="POST" action="/words/">
		Enter a term <input type="text" name="term">
		Enter its definition <textarea name="def"></textarea>
		<input type="submit">
	</form>`)
}

func getWord(res http.ResponseWriter, req *http.Request){
	ctx := appengine.NewContext(req)
	term := req.FormValue("term")


	q := datastore.NewQuery("Word").Order("Term").Filter("Term =", term)

	html := " "

	iterator := q.Run(ctx)

	for{
		var entity Word
		_, err := iterator.Next(&entity)
		if err == datastore.Done{
			break
		}else if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		html += `<dt>` +entity.Term+ `</dt>
				 <dd>` +entity.Def+ `</dd>`
	}
	

	res.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(res, `
	<dl>`+html+`<dl>
	<form method="POST" action="">
		Enter a term to see its definition <input type="text" name="term">
		<input type="submit">
	</form>
		`)
}
