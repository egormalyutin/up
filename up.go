package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	settings Settings
	files    http.Handler
)

/////////////////////////

func GETHandler(w http.ResponseWriter, r *http.Request) {
	// get link
	link := r.URL.String()[len("/api/"):]

	result := Check(link)
	fmt.Fprintf(w, result)

	if settings.Debug {
		log.Printf("%s: %s\n", link, result)
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case !settings.OnlyAPI && ((r.URL.Path == "/index.html") || (r.URL.Path == "/main.css") || (r.URL.Path == "/main.js") || (r.URL.Path == "/")):
		files.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/api"):
		GETHandler(w, r)
	default:
		http.Error(w, "404", http.StatusNotFound)
	}
}

/////////////////////////

func Serve() {
	settings = GetSettings()

	files = http.FileServer(assetFS())

	if !settings.OnlyAPI {
		files = http.FileServer(assetFS())
	}

	http.HandleFunc("/", AllHandler)

	log.Printf("Listening on http://localhost:%d\n", settings.Port)

	portString := fmt.Sprintf(":%d", settings.Port)
	log.Fatal(http.ListenAndServe(portString, nil))
}

/////////////////////////

func main() {
	Serve()
}
