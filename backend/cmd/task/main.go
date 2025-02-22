package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

const (
	targetDomain = "www.syllabus.kit.ac.jp"
	baseURL      = "https://" + targetDomain + "/"
)

func main() {
	page := 1

	c := colly.NewCollector(
		colly.AllowedDomains(targetDomain),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{Parallelism: 5})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("#search_result table", func(e *colly.HTMLElement) {
		if e.DOM.HasClass("data_list_tbl") {
			scrapeCourse(e)
		}
	})

	c.OnHTML("p.paging_area", func(e *colly.HTMLElement) {
		nextElem := e.DOM.Children().Last()
		text := nextElem.Text()
		if strings.Contains(text, "次へ") {
			href, _ := nextElem.Attr("href")
			targetURL := baseURL + href
			c.Visit(targetURL)
		}
	})

	targetURL := baseURL + fmt.Sprintf("?c=search_list&sk=&dc=01&page=%d", page)
	err := c.Visit(targetURL)
	if err != nil {
		log.Println(err)
	}
	c.Wait()
}

func scrapeCourse(e *colly.HTMLElement) {
	e.ForEach("tr", func(i int, row *colly.HTMLElement) {
		rowData := []string{}
		row.ForEach("td", func(i int, col *colly.HTMLElement) {
			data := ""
			if col.DOM.Find("form > a").Length() > 0 {
				data = col.DOM.Find("form > a").Contents().First().Text()
			} else {
				data = col.DOM.Contents().First().Text()
			}
			rowData = append(rowData, data)
		})
		for i := 0; i < len(rowData)-1; i++ {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(rowData[i])
		}

		if len(rowData) > 0 {
			fmt.Println()
		}
	})
}
