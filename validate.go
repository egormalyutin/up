package main

import (
	"regexp"

	cregex "github.com/mingrammer/commonregex"
)

var (
	protocolRG = regexp.MustCompile("^[\\d\\w]+:/+")
	hostRG     = regexp.MustCompile("^(.+)/.*$")
	portRG     = regexp.MustCompile(":\\d+$")

	privateRanges = [7]*regexp.Regexp{
		regexp.MustCompile("^(::f{4}:)?10\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}"),
		regexp.MustCompile("^(::f{4}:)?127\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}"),
		regexp.MustCompile("^(::f{4}:)?169\\.254\\.([1-9]|1?\\d\\d|2[0-4]\\d|25[0-4])\\.\\d{1,3}"),
		regexp.MustCompile("^(::f{4}:)?(172\\.1[6-9]|172\\.2\\d|172\\.3[0-1])\\.\\d{1,3}\\.\\d{1,3}"),
		regexp.MustCompile("^(::f{4}:)?192\\.168\\.\\d{1,3}\\.\\d{1,3}"),
		regexp.MustCompile("^f[c-d][0-9a-f]{2}(::1$|:[0-9a-f]{1,4}){1,7}"),
		regexp.MustCompile("^fe[89ab][0-9a-f](::1$|:[0-9a-f]{1,4}){1,7}"),
	}
)

func IsIP(link string) bool {
	return len(cregex.IPs(link)) > 0
}

func IsPrivate(link string) bool {
	if (link == "::") || (link == "::1") {
		return true
	} else {
		for _, rg := range privateRanges {
			if rg.FindString(link) != "" {
				return true
			}
		}
		return false
	}
}

// IsURL checks if the string is an URL.
func IsURL(str string) bool {
	return len(cregex.Links(str)) > 0
}

// Convert invalid URL to valid URL.
func ValidateURL(link string) string {
	// Remove protocol from link
	link = protocolRG.ReplaceAllString(link, "")

	// Remove path from link
	link = hostRG.ReplaceAllString(link, "$1")

	// Add port to link, if there are no another port
	portMatch := portRG.FindString(link)
	if portMatch == "" {
		link += ":80"
	}

	return link
}
