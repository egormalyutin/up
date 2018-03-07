package main

import (
	"net"
	"time"
)

const (
	DEFAULT_TIMEOUT = time.Duration(5 * time.Second)
)

func Fetch(link string) string {
	resp, err := net.DialTimeout("tcp", link, DEFAULT_TIMEOUT)
	if err == nil {
		resp.Close()
		return GenerateUpJSON()
	} else {
		return GenerateDownJSON()
	}
}

func Check(link string) string {
	link = ValidateURL(link)

	if IsURL(link) {
		if IsIP(link) {
			if IsPrivate(link) {
				return GenerateErrorJSON(1)
			} else {
				return Fetch(link)
			}
		} else {
			return Fetch(link)
		}
	} else {
		return GenerateErrorJSON(0)
	}
}
