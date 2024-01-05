package main

import (
	"net/http"
	"text/template"
)

//var store *sessions.FilesystemStore

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	// Save the session
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/index.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}
