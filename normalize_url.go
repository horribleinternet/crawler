package main

import (
	"net/url"
)

func normalizeURL(URL string) (string, error) {

	norm, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	path := norm.EscapedPath()
	count := 0
	for path[len(path)-1-count] == '/' {
		count++
	}
	path = path[:len(path)-count]
	return norm.Host + path, nil
}
