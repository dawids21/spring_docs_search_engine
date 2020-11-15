package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {

	fi, _ := os.Create("file.txt")

	urlsFile, _ := os.Open("urls.txt")

	var text []string

	scanner := bufio.NewScanner(urlsFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, s := range text {
		sub := strings.Split(s, "|")

		fmt.Println(strings.TrimSpace(sub[0]), strings.TrimSpace(sub[1]))
	}

	os.Exit(1)

	c := colly.NewCollector(
		colly.AllowedDomains("docs.spring.io"),
	)

	c.Limit(&colly.LimitRule{
		Delay:       3 * time.Second,
		RandomDelay: 1 * time.Second,
		DomainGlob:  "docs.spring.io/*",
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		fmt.Fprintf(fi, "%v=%v\n", h.Text, h.Request.AbsoluteURL(link))
	})

	c.Visit(url)
}
