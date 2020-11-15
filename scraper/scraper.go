package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {

	url := "https://docs.spring.io/spring-framework/docs/current/javadoc-api/allclasses-noframe.html"

	fi, _ := os.Create("file.txt")

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
