package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewFilesystemStore("/var/lib/access-www/session", []byte("super-secret-key"))

func main() {

	certFile := "/usr/share/access-www/ssl/access.baf.crt"
	keyFile := "/usr/share/access-www/ssl/access.baf.key"

	// create Gorilla router
	router := mux.NewRouter()

	// define the absolute path to the static files directory
	absolutePath, err := filepath.Abs("/usr/share/access-www/assets/")
	if err != nil {
		panic(err)
	}

	// create a file server for the static files
	fileServer := http.FileServer(http.Dir(absolutePath))

	// static files route
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	// home route
	router.HandleFunc("/", HomeHandler)

	// visitor in routes
	router.HandleFunc("/step-1-visitor-in", StepOneVisitorInHandler)
	router.HandleFunc("/step-2-visitor-in", StepTwoVisitorInHandler)
	router.HandleFunc("/step-3-visitor-in", StepThreeVisitorInHandler)
	router.HandleFunc("/step-4-visitor-in", StepFourVisitorInHandler)

	// visitor in api
	router.HandleFunc("/api/v1/step-1-visitor-in", ApiStepOneVisitorInHandler)
	router.HandleFunc("/api/v1/step-2-visitor-in", ApiStepTwoVisitorInHandler)
	router.HandleFunc("/api/v1/step-3-visitor-in", ApiStepThreeVisitorInHandler)
	router.HandleFunc("/api/v1/step-abort-visitor-in", ApiStepAbortVisitorInHandler)

	// make appointment route
	router.HandleFunc("/step-1-make-appointment", StepOneMakeAppointmentHandler)
	router.HandleFunc("/step-2-make-appointment", StepTwoMakeAppointmentHandler)

	// make appointment api
	router.HandleFunc("/api/v1/appointment-check-uuid/{uuid}", VerifyUUIDHandler)
	router.HandleFunc("/api/v1/step-1-make-appointment", ApiStepOneMakeAppointmentHandler)

	// visitor status route
	router.HandleFunc("/visitor-status", VisistorStatusHandler)
	router.HandleFunc("/api/v1/check-visitor-status", ApiCheckVisistorStatusHandler)

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

	// HTTP server
	fmt.Println("[+] http server listening at :80")
	log.Fatal(httpServer.ListenAndServe())
}
