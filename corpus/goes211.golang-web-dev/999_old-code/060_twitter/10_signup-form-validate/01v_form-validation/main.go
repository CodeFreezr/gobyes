package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"html/template"
	"io/ioutil"
	"log"
	stdlog "log"
	"net/http"
)

type User struct {
	Email    string
	UserName string
	Password string
}

var tpl *template.Template

func init() {
	r := httprouter.New()
	http.Handle("/", r)
	r.GET("/", Home)
	r.GET("/login", Login)
	r.GET("/signup", Signup)
	r.POST("/api/checkusername", checkUserName)
	r.POST("/api/createuser", createUser)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	tpl = template.Must(template.ParseGlob("templates/html/*.html"))
}

func Home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func Login(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "login.html", nil)
}

func Signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

func checkUserName(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	stdlog.Fatalln("REQUEST BODY: ", sbs)
	q, err := datastore.NewQuery("Users").Filter("UserName=", sbs).Count(ctx)
	stdlog.Fatalln("ERR: ", err)
	stdlog.Fatalln("QUANTITY: ", q)
	if err != nil {
		fmt.Fprint(res, "false")
		return
	}
	if q >= 1 {
		fmt.Fprint(res, "true")
	} else {
		fmt.Fprint(res, "false")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	NewUser := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("userName"),
		Password: req.FormValue("password"),
	}
	log.Fatalln(NewUser)
	fmt.Fprintln(res, NewUser.Email, NewUser.Password, NewUser.UserName)
}
