package main

import (
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {

	url := "https://docs.spring.io/spring-framework/docs/current/javadoc-api/allclasses-noframe.html"

	c := colly.NewCollector(
		colly.AllowedDomains("docs.spring.io"),
	)

	c.Limit(&colly.LimitRule{
		Delay:       3 * time.Second,
		RandomDelay: 1 * time.Second,
		DomainGlob:  "docs.spring.io/*",
	})

}
