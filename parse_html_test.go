package main

import "testing"

func TestGetH1FromHTMLBasic(t *testing.T) {
	tests := []struct {
		name     string
		inpuHTML string
		expected string
	}{
		{
			name:     "simple test",
			inpuHTML: "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "two h1 test",
			inpuHTML: "<html><body><h1>Test Title</h1><h2>smaller title</h2><h1>another Title</h1></body></html>",
			expected: "Test Title",
		},
		{
			name:     "no h1 test",
			inpuHTML: "<html><body><p>Outside paragraph.</p><main><p>Main paragraph.</p></main></body></html>",
			expected: "",
		},
		{
			name:     "complicated h1 test",
			inpuHTML: "<html><body><p>Outside paragraph.</p><main><h1>Test Title</h1><p>Main paragraph.</p><h1>another Title</h1></main></body></html>",
			expected: "Test Title",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getH1FromHTML(tc.inpuHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {

	tests := []struct {
		name     string
		inpuHTML string
		expected string
	}{
		{
			name:     "main section test",
			inpuHTML: "<html><body><p>Outside paragraph.</p><main><p>Main paragraph.</p></main></body></html>",
			expected: "Main paragraph.",
		},
		{
			name:     "empty main section test",
			inpuHTML: "<html><body><p>Outside paragraph.</p><main></main></body></html>",
			expected: "Outside paragraph.",
		},
		{
			name:     "no paragaraph test",
			inpuHTML: "<html><body><main></main></body></html>",
			expected: "",
		},
		{
			name:     "no main section test",
			inpuHTML: "<html><body><p>Outside paragraph.</p></body></html>",
			expected: "Outside paragraph.",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inpuHTML)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
