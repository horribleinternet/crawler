package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTMLAbsolute(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "basic test",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="https://blog.boot.dev"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://blog.boot.dev"},
		},
		{
			name:      "relative test",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a href="docs/index.html"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://blog.boot.dev/docs/index.html"},
		},
		{
			name:      "multi test",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><a referrerpolicy="no-referrer" href="docs/index.html"><span>Boot.dev</span></a><p>Some text.</p><a href="https://www.boot.dev"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://blog.boot.dev/docs/index.html", "https://www.boot.dev"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - %s FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}

func TestGetImagesFromHTMLRelative(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "basic test",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png"},
		},
		{
			name:      "multi test",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"><p>Some text.</p><img alt="Cat" src="/cat.png"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png", "https://blog.boot.dev/cat.png"},
		},
		{
			name:      "absolute test",
			inputURL:  "https://blog.boot.dev",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"><p>Some text.</p><img alt="Cat" src="https://www.boot.dev/cat.png"></body></html>`,
			expected:  []string{"https://blog.boot.dev/logo.png", "https://www.boot.dev/cat.png"},
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - %s FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}

			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("Test %v - %s FAIL: unexpected error: %v", i, tc.name, err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, got %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
