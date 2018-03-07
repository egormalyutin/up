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
