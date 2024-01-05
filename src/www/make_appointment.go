package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func StepOneMakeAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["appointment"] = "step-0"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/make-appointment-1.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiStepOneMakeAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Read the request body
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "Error reading request body", http.StatusInternalServerError)
	//	return
	///}

	// Close the request body
	//defer r.Body.Close()

	///filename := "/var/lib/access-www/session/visitor_data_" + session.ID
	///err = ioutil.WriteFile(filename, []byte(body), 0644)

	//if err != nil {
	///fmt.Println("Error writing to file:", err)
	//	http.Error(w, "Error writing to file", http.StatusInternalServerError)
	//	return
	//}

	session.Values["appointment"] = "step-1"
	session.Save(r, w)

	uuid4, _ := uuid.NewRandom()

	// send the request to the admin microservice

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\":true, \"next\": \"/step-2-make-appointment\", \"request-number\":\"" + uuid4.String() + "\"}"))

}

func StepTwoMakeAppointmentHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/make-appointment-2.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func VerifyUUIDHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")
	session.Save(r, w)

	vars := mux.Vars(r)

	fmt.Printf("UUID: %s\n", vars["uuid"])

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\":true}"))
}
