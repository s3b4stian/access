package main

import (
	"net/http"
	"text/template"
)

func VisistorStatusHandler(w http.ResponseWriter, r *http.Request) {
	//session, _ := store.Get(r, "access-session")

	//session.Values["appointment"] = "step-0"
	//session.Save(r, w)

	tmpl := template.Must(template.ParseFiles("/usr/share/access-www/templates/visitor-status.html", "/usr/share/access-www/templates/header.html", "/usr/share/access-www/templates/footer.html"))
	tmpl.Execute(w, nil)
}

func ApiCheckVisistorStatusHandler(w http.ResponseWriter, r *http.Request) {

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

	// set the Content-Type header to application/json and return response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"success\":true, \"data\"}"))

}
