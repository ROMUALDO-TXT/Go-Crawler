package crawler

import (
	"GoCrawler/model"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Scrap(urlToProcess []string, rchan chan model.Result) {
	defer close(rchan)

	var results = []chan model.Result{}

	for i, url := range urlToProcess {
		results = append(results, make(chan model.Result))
		go scrapParallel(url, results[i])
	}

	for j := range results {
		for r1 := range results[j] {
			rchan <- r1
		}
	}

}

func scrapParallel(url string, rchan chan model.Result) {
	defer close(rchan)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: It can't scrap '", url, "'")
	}

	defer resp.Body.Close()
	body := resp.Body
	htmlParsed, err := html.Parse(body)
	if err != nil {
		fmt.Println("ERROR: It can't parse html '", url, "'")
	}
	header := GetFirstElementByClass(htmlParsed, "header", "")
	var r model.Result
	a := GetFirstElementByClass(header, "a", "ds-link--styleSubtle")
	r.UserName = GetFirstTextNode(a).Data

	div := GetFirstElementByClass(htmlParsed, "div", "section-content")
	h1 := GetFirstElementByClass(div, "h1", "graf--title")
	r.Title = GetFirstTextNode(h1).Data

	footer := GetFirstElementByClass(htmlParsed, "footer", "u-paddingTop10")
	buttonLikes := GetFirstElementByClass(footer, "button", "js-multirecommendCountButton")
	r.Likes = GetFirstTextNode(buttonLikes).Data

	rchan <- r

}
