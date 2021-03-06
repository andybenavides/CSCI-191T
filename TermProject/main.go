package main

import(
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type User struct{
	FirstName string
	LastName string
	UserName string
	Email string
	Password string
}

type SessionData struct{
	User
	LoggedIn bool
	LogInFail bool
}

var tpl *template.Template

func init(){
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/login", Login)
	r.GET("/signup", SignUp)
	r.POST("/form/createuser", CreateUser)
	r.POST("/form/login", UserLogin)
	r.GET("/logout", LogOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	tpl.ExecuteTemplate(res, "homePage.html", nil)
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	tpl.ExecuteTemplate(res, "loginPage.html", nil)
}

func SignUp(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	tpl.ExecuteTemplate(res, "signUp.html", nil)
}

func LogOut(res http.ResponseWriter, req *http.Request, _ httprouter.Params){
	var sd SessionData
	sd.LoggedIn = false
	tpl.ExecuteTemplate(res, "homePage.html", &sd)
}