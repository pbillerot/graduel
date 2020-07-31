package views

/*Holds the fetch related view handlers*/

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/csrf"
	"github.com/pbillerot/graduel/config"
	"github.com/pbillerot/graduel/dico"
	"github.com/pbillerot/graduel/types"
)

var homeTemplate *template.Template
var aboutTemplate *template.Template
var templates *template.Template
var loginTemplate *template.Template

var message string //message will store the message to be shown as notification
var err error
var application dico.Application

//ShowPortailFunc is used to handle the "/" URL which is the default ons
func ShowPortailFunc(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] != true {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method == "GET" {
		message = "Hello !!!"
		if message != "" {
			context.Set(r, "Message", message)
		}
		GraduelAddContext(r)
		context.Set(r, "Application", dico.GetDico())
		homeTemplate.Execute(w, context.GetAll(r))
	}
}

// AboutFunc as
func AboutFunc(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] != true {
		http.Redirect(w, r, "/login", 302)
		return
	}
	configFile, err := ioutil.ReadFile("about.json")
	if err != nil {
		log.Print("Unable to read about.json file")
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}
	//log.Print(configFile)
	About := types.About{}
	err = json.Unmarshal(configFile, &About)
	if err != nil {
		log.Print("Invalid about.json")
		http.Redirect(w, r, "/", http.StatusInternalServerError)
	}

	GraduelAddContext(r)
	context.Set(r, "About", About)
	context.Set(r, "Application", dico.GetDico())
	aboutTemplate.Execute(w, context.GetAll(r))
}

//PopulateTemplates is used to parse all templates present in the templates folder
func PopulateTemplates() {
	var allFiles []string
	templatesDir := "./templates/"
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1) // No point in running app if templates aren't read
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	homeTemplate = templates.Lookup("home.html")
	aboutTemplate = templates.Lookup("about.html")
	loginTemplate = templates.Lookup("login.html")

}

// GraduelAddContext ajout dans le context des données de sessions, config, application
func GraduelAddContext(r *http.Request) {
	// Ajout des données de session
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] == true {
		context.Set(r, "loggedin", true)
		context.Set(r, "username", session.Values["username"])
	}
	// Ajout de config
	conf, _ := config.ReadConfig()
	context.Set(r, "config", conf)
	// Ajout du token de sécurité
	context.Set(r, "CSRFToken", csrf.TemplateField(r))
	// Ajout du dictionnaire de l'application
	context.Set(r, "Application", dico.GetDico())
}
