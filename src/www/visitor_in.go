package main

import (
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/google/uuid"
)

func StepOneVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["visitor"] = "step-0"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-1.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiStepOneVisitorInHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "access-session")

	// check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// close the request body
	defer r.Body.Close()

	filename := "/var/lib/access-www/session/visitor_data_" + session.ID
	err = ioutil.WriteFile(filename, []byte(body), 0644)

	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	session.Values["visitor"] = "step-1"
	session.Save(r, w)

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\":true, \"next\": \"/step-2-visitor-in\"}"))
}

func StepTwoVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["visitor"] = "step-2"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-2.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiStepTwoVisitorInHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "access-session")

	// check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// close the request body
	defer r.Body.Close()

	filename := "/var/lib/access-www/session/visitor_photo_" + session.ID
	err = ioutil.WriteFile(filename, []byte(body), 0644)

	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	session.Values["visitor"] = "step-3"
	session.Save(r, w)

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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

	// check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// eead the request body
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, "Error reading request body", http.StatusInternalServerError)
	//	return
	//}

	// close the request body
	//defer r.Body.Close()

	//fmt.Printf("Request Body DATA: %s... and more\n", body)

	session.Values["visitor"] = "step-5"
	session.Save(r, w)

	// generate a new uuidv4
	uuid4, _ := uuid.NewRandom()

	// this part is no more mandatory after the backend go up
	filename := "/var/lib/access-www/session/visitor_uuid_" + session.ID
	err := ioutil.WriteFile(filename, []byte(uuid4.String()), 0644)

	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	// send the request to the admin microservice

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
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

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\":true, \"next\": \"/\"}"))
}

func StepFourVisitorInHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "access-session")

	session.Values["visitor"] = "step-6"
	session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-in-4.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}
