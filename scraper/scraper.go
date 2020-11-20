package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	lines := make([]string, 0)

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
		nextLine := fmt.Sprintf("%v=%v", h.Text, h.Request.AbsoluteURL(link))
		lines = append(lines, nextLine)
	})

	for name, url := range urls {
		err := c.Visit(url)

		if err == nil {

			fi, _ := os.Create("classes_urls/" + name + ".txt")
			sort.Strings(lines)

			for _, line := range lines {
				fmt.Fprintf(fi, "%v\n", line)
			}
		} else {
			os.Exit(1)
		}
	}
}
