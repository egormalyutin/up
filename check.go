package main

import (
	"net"
	"time"
)

const (
	DEFAULT_TIMEOUT = time.Duration(800 * time.Millisecond)
)

func Check(link string) string {
	link = ValidateURL(link)

	if IsURL(link) {
		resp, err := net.DialTimeout("tcp", link, DEFAULT_TIMEOUT)
		if err == nil {
			resp.Close()
			return "up"
		} else {
			return "down"
		}
	} else {
		return "error"
	}
}
