package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	start        time.Time
	isArchive    bool
	year         int
	page         int
	targetDomain string
	baseURL      string
)

func init() {
	start = time.Now()
	flag.BoolVar(&isArchive, "archive", false, "whether to retrieve info from the archive")
	flag.IntVar(&year, "year", start.Year(), "archive year")
	flag.IntVar(&page, "page", 1, "page offset")
	flag.Parse()

	targetDomain = "www.syllabus.kit.ac.jp"
	baseURL = "https://" + targetDomain + "/"
	if isArchive {
		baseURL += "archive/"
	}
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(targetDomain),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{Parallelism: 2})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML("#search_result table.data_list_tbl", func(e *colly.HTMLElement) {
		scrapeCourse(e)
	})

	c.OnHTML("p.paging_area a[href]", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "次へ") {
			href := e.Attr("href")
			_ = e.Request.Visit(baseURL + href)
		}
	})

	param := "?c=search_list&sk=&dc=01" + fmt.Sprintf("&yr=%d&page=%d", year, page)
	targetURL := baseURL + param

	err := c.Visit(targetURL)
	if err != nil {
		log.Println(err)
	}

	c.Wait()

	log.Println("Total time:", time.Since(start))
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
