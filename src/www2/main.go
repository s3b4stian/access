package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("super-secret-key"))

func main() {

	certFile := "/usr/share/access-www/ssl/access.baf.crt"
	keyFile := "/usr/share/access-www/ssl/access.baf.key"

	// Create Gorilla router
	router := mux.NewRouter()

	// Define the absolute path to the static files directory
	absolutePath, err := filepath.Abs("/usr/share/access-www/assets/")
	if err != nil {
		panic(err)
	}

	// Create a file server for the static files
	fileServer := http.FileServer(http.Dir(absolutePath))

	// Handle static files using Gorilla Mux
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/visitor-in", VisitorInHandler)
	router.HandleFunc("/step-1-visitor-in", StepOneVisitorInHandler)
	router.HandleFunc("/step-2-visitor-in", StepTwoVisitorInHandler)
	router.HandleFunc("/step-3-visitor-in", StepThreeVisitorInHandler)
	router.HandleFunc("/step-4-visitor-in", StepFourVisitorInHandler)

	httpServer := &http.Server{
		Handler:      router,
		Addr:         ":80",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	tlsServer := &http.Server{
		Handler:      router,
		Addr:         ":443",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
		TLSConfig:    nil,
	}

	// HTTPS server
	go func() {
		fmt.Println("[+] https server listening at :443")
		log.Fatal(tlsServer.ListenAndServeTLS(certFile, keyFile))

	}()

	fmt.Println("[+] http server listening at :80")
	log.Fatal(httpServer.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	// Set session values
	//session.Values["username"] = "john_doe"
	//session.Values["authenticated"] = true

	// Save the session
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/index.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func VisitorInHandler(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "access-session")
}

func StepOneVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "access-session")

	// Retrieve session values
	//username, ok := session.Values["username"].(string)
	//authenticated, _ := session.Values["authenticated"].(bool)

	//if !ok || !authenticated {
	//	w.Write([]byte("Not authenticated "))
	//} else {
	//	w.Write([]byte("Authenticated user: " + username + " "))
	//}

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-1.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func StepTwoVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-2.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func StepThreeVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-3.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func StepFourVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-4.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}
