package main

import (
	"os"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
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

	q, err := queue.New(
		4,
		&queue.InMemoryQueueStorage{MaxSize: 1000},
	)

	if err != nil {
		os.Exit(1)
	}

	q.AddURL(url)

	q.Run(c)
}
