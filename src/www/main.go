package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

//var store = sessions.NewCookieStore([]byte("super-secret-key"))
var store = sessions.NewFilesystemStore("/tmp", []byte("super-secret-key"))

func main() {

	//store.MaxLength(4096 * 3)

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
	router.HandleFunc("/step-1-visitor-in", StepOneVisitorInHandler)
	router.HandleFunc("/step-2-visitor-in", StepTwoVisitorInHandler)
	router.HandleFunc("/step-3-visitor-in", StepThreeVisitorInHandler)
	router.HandleFunc("/step-4-visitor-in", StepFourVisitorInHandler)

	router.HandleFunc("/api/v1/step-1-visitor-in", ApiStepOneVisitorInHandler)
	router.HandleFunc("/api/v1/step-2-visitor-in", ApiStepTwoVisitorInHandler)
	router.HandleFunc("/api/v1/step-3-visitor-in", ApiStepThreeVisitorInHandler)
	router.HandleFunc("/api/v1/step-abort-visitor-in", ApiStepAbortVisitorInHandler)

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

	// Save the session
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/index.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func StepOneVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["visitor"] = "step-0"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-1.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiStepOneVisitorInHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "access-session")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Close the request body
	defer r.Body.Close()

	filename := "/tmp/visitor_data_" + session.ID
	err = ioutil.WriteFile(filename, []byte(body), 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	session.Values["visitor"] = "step-1"
	session.Save(r, w)

	// Write the JSON data to the response body
	w.Write([]byte("{\"success\":true, \"next\": \"/step-2-visitor-in\"}"))

}

func StepTwoVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	//visitor_data, _ := session.Values["visitor"].(string)

	//fmt.Printf("Session ID: %s\n", session.ID)
	//fmt.Printf("Session DATA: %s\n", visitor_data)

	session.Values["visitor"] = "step-2"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-2.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiStepTwoVisitorInHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "access-session")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Close the request body
	defer r.Body.Close()

	filename := "/tmp/visitor_photo_" + session.ID
	err = ioutil.WriteFile(filename, []byte(body), 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	session.Values["visitor"] = "step-3"
	session.Save(r, w)

	// Write the JSON data to the response body
	w.Write([]byte("{\"success\":true, \"next\": \"/step-3-visitor-in\"}"))

}

func StepThreeVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["visitor"] = "step-4"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-3.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiStepThreeVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "Error reading request body", http.StatusInternalServerError)
	//	return
	//}

	// Close the request body
	//defer r.Body.Close()

	//fmt.Printf("Request Body DATA: %s... and more\n", body)

	session.Values["visitor"] = "step-5"
	session.Save(r, w)

	uuid4, _ := uuid.NewRandom()

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	filename := "/tmp/visitor_uuid_" + session.ID
	err := ioutil.WriteFile(filename, []byte(uuid4.String()), 0644)

	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Write the JSON data to the response body
	w.Write([]byte("{\"success\":true, \"next\": \"/step-4-visitor-in\", \"request-number\":\"" + uuid4.String() + "\"}"))

}

func ApiStepAbortVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	session.Values["visitor"] = "step-0"
	session.Save(r, w)

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response body
	w.Write([]byte("{\"success\":true, \"next\": \"/\"}"))
}

func StepFourVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["visitor"] = "step-6"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-4.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}
