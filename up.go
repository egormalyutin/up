package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	DEFAULT_TIMEOUT = time.Duration(6000 * time.Millisecond)
)

type Settings struct {
	Port  int
	Debug bool
}

var settings Settings

func GetSettings() Settings {
	var port int64
	portEnv := os.Getenv("PORT")
	if portEnv != "" {
		var err error
		port, err = strconv.ParseInt(portEnv, 10, 64)
		if err != nil {
			log.Fatal("Wrong env variable PORT!")
		}
	} else {
		port = 8080 // default
	}

	var debug bool
	debugEnv := os.Getenv("DEBUG")
	if debugEnv != "" {
		debug = true
	} else {
		debug = false // default
	}

	settings := Settings{
		Port:  int(port),
		Debug: debug,
	}

	return settings
}

/////////////////////////

func GenerateUpJSON(status int) string {
	mp := map[string]interface{}{"ok": true, "status": status}
	j, err := json.Marshal(&mp)
	if err != nil {
		return ""
	} else {
		return string(j)
	}
}

func GenerateDownJSON() string {
	mp := map[string]bool{"ok": false}
	j, err := json.Marshal(&mp)
	if err != nil {
		return ""
	} else {
		return string(j)
	}
}

/////////////////////////

func ValidateURL(link string) string {
	u, err := url.Parse(link)

	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "http"

	rg := regexp.MustCompile("^/*")
	u.Path = rg.ReplaceAllString(u.Path, "")

	return u.String()
}

/////////////////////////

func GETHandler(w http.ResponseWriter, r *http.Request) {
	// get link
	link := r.URL.Path[len("/api/"):]

	link = ValidateURL(link)

	// create http client
	client := http.Client{
		Timeout: DEFAULT_TIMEOUT,
	}

	// GET link
	resp, err2 := client.Get(link)

	var result bool

	if err2 == nil {
		result = true
	} else {
		result = false
	}

	var resultString, j string
	if result {
		resultString = "up"
		j = GenerateUpJSON(resp.StatusCode)
	} else {
		resultString = "down"
		j = GenerateDownJSON()
	}

	fmt.Fprintf(w, j)

	if settings.Debug {
		log.Printf("%s: %s\n", link, resultString)
	}
}

/////////////////////////

func main() {
	settings = GetSettings()

	http.Handle("/", http.FileServer(assetFS()))
	http.HandleFunc("/api/", GETHandler)

	log.Printf("Listening on http://localhost:%d\n", settings.Port)

	portString := fmt.Sprintf(":%d", settings.Port)

	log.Fatal(http.ListenAndServe(portString, nil))
}
