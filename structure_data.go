package main

import "net/url"

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	pg := PageData{
		URL:            pageURL,
		H1:             getH1FromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
	}
	url, err := url.Parse(pageURL)
	if err == nil {
		pg.OutgoingLinks, _ = getURLsFromHTML(html, url)
		pg.ImageURLs, _ = getImagesFromHTML(html, url)
	}
	return pg
}
