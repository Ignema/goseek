package main

import (
	"fmt"
	"io"
	"strings"
	"syscall/js"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Attaching golang function to context...")
	js.Global().Set("exec_worker", js.FuncOf(scrape))
	select {}
}

func scrape(this js.Value, args []js.Value) interface{}  {
	doc, err := goquery.NewDocumentFromReader(io.NopCloser(strings.NewReader(args[0].String())))
	if err != nil {
		panic(err)
	}

	var results []string
	doc.Find(args[1].String()).Each(func(i int, s *goquery.Selection) {
		html, err := s.Html()
		if err != nil {
			panic(err)
		}
		results[i] = html
	})

	return js.ValueOf(results)
}