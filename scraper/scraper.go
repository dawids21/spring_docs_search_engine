package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func getUrlsFromFile(fileName string) map[string]string {

	urlsFile, _ := os.Open(fileName)

	var text []string

	scanner := bufio.NewScanner(urlsFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	urls := make(map[string]string)

	for _, s := range text {
		sub := strings.Split(s, "|")
		urls[sub[0]] = sub[1]
	}

	return urls
}

func main() {

	fi, _ := os.Create("file.txt")

	urls := getUrlsFromFile("urls.txt")

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
