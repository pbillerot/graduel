package views

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/pbillerot/graduel/db"
	"github.com/pbillerot/graduel/dico"
)

// CodeSecure as
var CodeSecure = []byte(os.Getenv("SESSION_KEY"))

// Store as
var Store = sessions.NewCookieStore(CodeSecure)

//RequiresLogin is a middleware which will be used for each httpHandler to check if there is any active session
func RequiresLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		if session.Values["loggedin"] != "true" {
			http.Redirect(w, r, "/login", 302)
			return
		}
		handler(w, r)
	}
}

//LogoutFunc Implements the logout functionality. WIll delete the session information from the cookie store
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err == nil { //If there is no error, then remove session
		if session.Values["loggedin"] != "false" {
			session.Values["loggedin"] = "false"
			session.Save(r, w)
		}
	}
	http.Redirect(w, r, "/login", 302) //redirect to login irrespective of error or not
}

//LoginFunc implements the login functionality, will add a cookie to the cookie store for managing authentication
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	context.Set(r, "CSRFToken", csrf.TemplateField(r))
	context.Set(r, "Application", dico.GetDico())

	switch r.Method {
	case "GET":
		log.Println("LoginFunc context", context.GetAll(r))
		loginTemplate.Execute(w, context.GetAll(r))
	case "POST":
		log.Print("Inside POST")
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if (username != "" && password != "") && db.ValidUser(username, password) {
			session.Values["loggedin"] = "true"
			session.Values["username"] = username
			session.Save(r, w)
			log.Print("user ", username, " is authenticated")
			context.Set(r, "username", username)
			http.Redirect(w, r, "/", 302)
			return
		}
		log.Print("Invalid user " + username)
		loginTemplate.Execute(w, context.GetAll(r))
	default:
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	}
}

//SignUpFunc will enable new users to sign up to our service
func SignUpFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")

	log.Println(username, password, email)

	err := db.CreateUser(username, password, email)
	if err != nil {
		http.Error(w, "Unable to sign user up", http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, "/login/", 302)
	}
}
