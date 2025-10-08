package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	inputURL := "https://blog.boot.dev"
	inputBody := `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
    </body></html>`

	actual := extractPageData(inputBody, inputURL)

	expected := PageData{
		URL:            "https://blog.boot.dev",
		H1:             "Test Title",
		FirstParagraph: "This is the first paragraph.",
		OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
		ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}

	inputURL = "https://blog.boot.dev"
	inputBody = `<html><body>
        <h1>Test Title</h1>
        <p>This is the first paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
        <p>This is the second paragraph.</p>
        <img src="/image2.jpg" alt="Image 1">
        <h1>Another Test Title</h1>
        <a href="https://www.boot.dev/link1">Link 1</a>
    </body></html>`

	actual = extractPageData(inputBody, inputURL)

	expected = PageData{
		URL:            "https://blog.boot.dev",
		H1:             "Test Title",
		FirstParagraph: "This is the first paragraph.",
		OutgoingLinks:  []string{"https://blog.boot.dev/link1", "https://www.boot.dev/link1"},
		ImageURLs:      []string{"https://blog.boot.dev/image1.jpg", "https://blog.boot.dev/image2.jpg"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}

	inputURL = "https://blog.boot.dev"
	inputBody = `<html><body>
        <h1>Stupid Title</h1>
        <p>This is a stupid paragraph.</p>
        <a href="/link1">Link 1</a>
        <img src="/image1.jpg" alt="Image 1">
        <p>This is the second paragraph.</p>
        <img src="/image2.jpg" alt="Image 1">
        <h1>Another Test Title</h1>
        <a href="https://www.boot.dev/link1">Link 1</a>
    </body></html>`

	actual = extractPageData(inputBody, inputURL)

	expected = PageData{
		URL:            "https://blog.boot.dev",
		H1:             "Stupid Title",
		FirstParagraph: "This is a stupid paragraph.",
		OutgoingLinks:  []string{"https://blog.boot.dev/link1", "https://www.boot.dev/link1"},
		ImageURLs:      []string{"https://blog.boot.dev/image1.jpg", "https://blog.boot.dev/image2.jpg"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
