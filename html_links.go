package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return getISelectorAttrFromHTML(htmlBody, baseURL, "a[href]", "href")
	/*
	   out := make([]string, 0)
	   doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))

	   	if err != nil {
	   		return out, err
	   	}

	   	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
	   		if len(s.Nodes) == 0 {
	   			return
	   		}
	   		for _, a := range s.Nodes[0].Attr {
	   			if a.Key == "href" {
	   				link, err := url.Parse(a.Val)
	   				if err != nil {
	   					return
	   				}
	   				link = baseURL.ResolveReference(link)
	   				out = append(out, link.String())
	   				return
	   			}
	   		}
	   	})

	   return out, nil
	*/
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	return getISelectorAttrFromHTML(htmlBody, baseURL, "img", "src")
	/*
		out := make([]string, 0)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
		if err != nil {
			return out, err
		}
		doc.Find("img").Each(func(_ int, s *goquery.Selection) {
			if len(s.Nodes) == 0 {
				return
			}
			for _, a := range s.Nodes[0].Attr {
				if a.Key == "src" {
					link, err := url.Parse(a.Val)
					if err != nil {
						return
					}
					fmt.Println(a.Val, link.String(), baseURL.String())
					link = baseURL.ResolveReference(link)
					fmt.Println(link.String())
					out = append(out, link.String())
					return
				}
			}
		})
		return out, nil
	*/
}

func getISelectorAttrFromHTML(htmlBody string, baseURL *url.URL, selector, attr string) ([]string, error) {
	out := make([]string, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return out, err
	}
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		if len(s.Nodes) == 0 {
			return
		}
		for _, a := range s.Nodes[0].Attr {
			if a.Key == attr {
				link, err := url.Parse(a.Val)
				if err != nil {
					return
				}
				fmt.Println(a.Val, link.String(), baseURL.String())
				link = baseURL.ResolveReference(link)
				fmt.Println(link.String())
				out = append(out, link.String())
				return
			}
		}
	})
	return out, nil
}
