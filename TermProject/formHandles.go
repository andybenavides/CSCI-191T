package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/golang/appengine"
	"github.com/golang/appengine/datastore"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	ctx := appengine.NewContext(req)

// Form validation skeleton
//
//	var PassVerify = req.FormValue("passWordVerify")
//	var Password = req.FormValue("passWord")
//
//	if(PassVerify != Password){
//
//	}

	var password, _ = bcrypt.GenerateFromPassword([]byte(req.FormValue("passWord")), 0)
	user := User{
		FirstName: req.FormValue("firstName"),
		LastName: req.FormValue("lastName"),
		UserName: req.FormValue("userName"),
		Email: req.FormValue("email"),
		Password: string(password),
	}

	key := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	key, err := datastore.Put(ctx,key, &user)
	if err != nil{
		http.Error(res, err.Error(), 500)
		return
	}

	tpl.ExecuteTemplate(res, "homePage.html", &user)
}

func UserLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	ctx := appengine.NewContext(req)

	key := datastore.NewKey(ctx, "Users", req.FormValue("userName"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if (err != nil) {
		var sd SessionData
		sd.LogInFail = true
		tpl.ExecuteTemplate(res, "loginPage.html", sd)
		return
	}else {
		user.UserName = req.FormValue("userName")
		tpl.ExecuteTemplate(res, "homePage.html", &user)
	}
}
