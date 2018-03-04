package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	DEFAULT_TIMEOUT = time.Duration(800 * time.Millisecond)

	RESULT_UP    = "up"
	RESULT_DOWN  = "down"
	RESULT_ERROR = "error"
)

var (
	settings Settings
	files    http.Handler
)

/////////////////////////

func GETHandler(w http.ResponseWriter, r *http.Request) {
	// get link
	link := r.URL.String()[len("/api/"):]

	link = ValidateURL(link)

	if IsURL(link) {
		// GET link
		resp, err2 := net.DialTimeout("tcp", link, DEFAULT_TIMEOUT)

		var result bool

		if err2 == nil {
			resp.Close()
			result = true
		} else {
			result = false
		}

		var resultString string

		if result {
			resultString = RESULT_UP
		} else {
			resultString = RESULT_DOWN
		}

		fmt.Fprintf(w, resultString)

		if settings.Debug {
			log.Printf("%s: %s\n", link, resultString)
		}
	} else {
		resultString := RESULT_ERROR
		fmt.Fprintf(w, resultString)
		if settings.Debug {
			log.Printf("%s: %s\n", link, resultString)
		}
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case !settings.OnlyAPI && ((strings.HasPrefix(r.URL.Path, "/index.html")) || (strings.HasPrefix(r.URL.Path, "/main.css")) || (strings.HasPrefix(r.URL.Path, "/main.js")) || (r.URL.Path == "/")):
		files.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/api"):
		GETHandler(w, r)
	default:
		http.Error(w, "404", http.StatusNotFound)
	}
}

/////////////////////////

func main() {
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
