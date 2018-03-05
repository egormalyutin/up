package main

import (
	"fmt"
)

var (
	errors = [2]string{
		"URL is invalid",
		"IP is private",
	}
)

// We can just use encoding/json...
// But should we?

func GenerateUpJSON() string {
	return "{\"up\":true}"
}

func GenerateDownJSON() string {
	return "{\"up\":false}"
}

func GenerateErrorJSON(code int) string {
	return fmt.Sprintf("{\"error\":\"%s\",\"errorCode\":%d}", errors[code], code)
}
