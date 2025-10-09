package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	html, err := getHTML(args[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(html)
	fmt.Printf("Starting crawl of: %s\n", args[0])
	pages := make(map[string]int)
	crawlPage(args[0], args[0], pages)
	for k, v := range pages {
		fmt.Printf("%70s : %d\n", k, v)
	}
}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Error parsing URL %s : %v\n", rawBaseURL, err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error parsing URL %s : %v\n", rawCurrentURL, err)
		return
	}
	if baseURL.Host != currentURL.Host {
		fmt.Printf("Not crawling %s\n", rawCurrentURL)
		return
	}
	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot normalize %s: %v\n", rawCurrentURL, err)
		return
	}
	if num, ok := pages[normURL]; ok {
		pages[normURL] = num + 1
		fmt.Printf("Already crawled %s\n", rawCurrentURL)
		return
	}
	pages[normURL] = 1
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot access %s: %v\n", rawCurrentURL, err)
	}
	uRLs, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("Cannot fetch links of %s: %v\n", rawCurrentURL, err)
	}
	fmt.Printf("Crawling %s\n", rawCurrentURL)
	for _, url := range uRLs {
		crawlPage(rawBaseURL, url, pages)
	}

}

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	req.Header.Add("User-Agent", "RootBeer/0.1")
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode < http.StatusBadRequest+100 {
		return "", fmt.Errorf("http error: %s", resp.Status)
	}
	if contype, ok := resp.Header[http.CanonicalHeaderKey("content-type")]; !ok || contype[0][:9] != "text/html" {
		return "", fmt.Errorf("not html content %s", contype[0])
	}
	defer resp.Body.Close()
	html, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(html), err
}
