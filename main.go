package main

import (
	"GoCrawler/crawler"
	"GoCrawler/model"
	"fmt"
	"time"
)

func main() {
	urlToProcess := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}

	initial := time.Now()

	r := make(chan model.Result)
	go crawler.Scrap(urlToProcess, r)

	for url := range r {
		fmt.Println(url)
	}

	fmt.Println("(Took ", time.Since(initial).Seconds(), "secs)")
}
