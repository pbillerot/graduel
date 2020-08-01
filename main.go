package main

/**
 * This is the main file for the Task application
 * License: MIT
 **/
import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/pbillerot/graduel/config"
	"github.com/pbillerot/graduel/dico"
	"github.com/pbillerot/graduel/views"
)

// Store as
var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	conf, err := config.ReadConfig()
	var port *string

	if err != nil {
		port = flag.String("port", "", "IP address")
		flag.Parse()

		//User is expected to give :8081 like input, if they give 8081
		//we'll append the required ':'
		if !strings.HasPrefix(*port, ":") {
			*port = ":" + *port
			log.Println("port is " + *port)
		}

		conf.ServerPort = *port
	}

	views.PopulateTemplates(conf.Template)
	r := mux.NewRouter()
	r.HandleFunc("/about", views.RequiresLogin(views.AboutFunc))
	r.HandleFunc("/login", views.LoginFunc)
	r.HandleFunc("/logout", views.RequiresLogin(views.LogoutFunc))
	r.HandleFunc("/signup", views.SignUpFunc)
	r.HandleFunc("/", views.RequiresLogin(views.ShowPortailFunc))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	r.HandleFunc("/favicon.ico", views.FaviconHandler)

	// http.HandleFunc("/api/get-task/", views.GetTasksFuncAPI)
	// http.HandleFunc("/api/get-deleted-task/", views.GetDeletedTaskFuncAPI)
	// http.HandleFunc("/api/add-task/", views.AddTaskFuncAPI)
	// http.HandleFunc("/api/update-task/", views.UpdateTaskFuncAPI)
	// http.HandleFunc("/api/delete-task/", views.DeleteTaskFuncAPI)

	// http.HandleFunc("/api/get-token/", views.GetTokenHandler)
	// http.HandleFunc("/api/get-category/", views.GetCategoryFuncAPI)
	// http.HandleFunc("/api/add-category/", views.AddCategoryFuncAPI)
	// http.HandleFunc("/api/update-category/", views.UpdateCategoryFuncAPI)
	// http.HandleFunc("/api/delete-category/", views.DeleteCategoryFuncAPI)

	// Chargement du dictionnaire
	// application, err := dico.LoadDico()
	// log.Println(application)

	dico.GetDico()

	log.Println(fmt.Printf("running server on http://localhost%s/login", conf.ServerPort))
	http.ListenAndServe(conf.ServerPort,
		wrapHandlerWithLogging(
			csrf.Protect(
				[]byte(os.Getenv("SECRET_KEY")),
				csrf.Secure(conf.CsrfSecure),
			)(r)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %d %s\n", r.RemoteAddr, r.Method, http.StatusOK, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// https://ndersson.me/post/capturing_status_code_in_net_http/
func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		lrw := newLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		log.Printf("%s %s %d %s", r.RemoteAddr, r.Method, statusCode, r.URL.Path)

		// log.Printf("<-- %d %s", statusCode, http.StatusText(statusCode))
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
