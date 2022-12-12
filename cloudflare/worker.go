package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	doc.Find(".titleline > a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			fmt.Printf(s.Text() + "\t" + link)
			} else {
			fmt.Printf("%s", s.Text())
		}
		fmt.Printf("\n")
	})
}