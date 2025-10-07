package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	selec := doc.Find("h1").First()
	return getChildData(selec)
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	selec := doc.Find("main").First()
	if selec != nil {
		selec = selec.Find("p").First()
		out := getChildData(selec)
		if len(out) > 0 {
			return out
		}
	}
	selec = doc.Find("p").First()
	return getChildData(selec)
}

func getChildData(selec *goquery.Selection) string {
	if selec != nil && len(selec.Nodes) > 0 && selec.Nodes[0].FirstChild != nil {
		return selec.Nodes[0].FirstChild.Data
	}
	return ""
}
