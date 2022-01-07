package internal

import (
	"log"
	"net/http"
)

// api routes with stripprefix and fileserver handler. StripPrefix and FileServer functions used for serve css styles for front-end pages.
func HandleRequests() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	ApiRoot := "/api"

	// .../api/homepage
	http.Handle(ApiRoot+"/homepage", IsAuthorized(Homepage))
	// .../api/register
	http.HandleFunc(ApiRoot+"/signup", SignUp)
	// .../api/login
	http.HandleFunc(ApiRoot+"/signin", SignIn)

	log.Fatal(http.ListenAndServe(":6363", nil))
}
