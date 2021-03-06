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

	var password, _ = bcrypt.GenerateFromPassword([]byte(req.FormValue("passWord")), bcrypt.DefaultCost)
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

	var sd SessionData
	sd.LoggedIn = true

	tpl.ExecuteTemplate(res, "homePage.html", &sd)
}

func UserLogin(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	ctx := appengine.NewContext(req)
	var sd SessionData

	key := datastore.NewKey(ctx, "Users", req.FormValue("userName"), 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	if (err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.FormValue("passWord"))) != nil) {
		sd.LogInFail = true
		tpl.ExecuteTemplate(res, "loginPage.html", sd)
		return
	}else {
		sd.LoggedIn = true
		sd.UserName = user.UserName
		tpl.ExecuteTemplate(res, "homePage.html", &sd)
	}
}


