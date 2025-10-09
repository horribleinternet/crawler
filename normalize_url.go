package main

import (
	"net/url"
	"strings"
)

func normalizeURL(URL string) (string, error) {

	norm, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	path := strings.TrimRight(norm.EscapedPath(), "/")
	return norm.Host + path, nil
}
