package task

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
	domain "github.com/yummynaan/kit-enrollment-helper/internal/domain"
	"github.com/yummynaan/kit-enrollment-helper/internal/domain/model"
)

type FetchSyllabusWorker struct {
	targetURL  *url.URL
	repository domain.Repository
}

func NewFetchSyllabusWorker(targetURL *url.URL, repository domain.Repository) *FetchSyllabusWorker {
	return &FetchSyllabusWorker{
		targetURL:  targetURL,
		repository: repository,
	}
}

func (w *FetchSyllabusWorker) Run() error {
	baseURL := fmt.Sprintf("%s://%s/", w.targetURL.Scheme, w.targetURL.Host)

	// colly, goquery, rodを用いてスクレイピング
	c := colly.NewCollector(
		colly.AllowedDomains(w.targetURL.Host),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{Parallelism: 10})

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

	err := c.Visit(w.targetURL.String())
	if err != nil {
		return err
	}

	c.Wait()

	return nil
}

func scrapeCourse(e *colly.HTMLElement) model.Courses {
	result := model.Courses{}
	isHeader := true
	e.ForEach("tr", func(i int, row *colly.HTMLElement) {
		if isHeader {
			isHeader = false
			return
		}
		count := 0
		rowData := []string{}
		row.ForEach("td", func(i int, col *colly.HTMLElement) {
			if count >= 9 {
				return
			}

			data := ""
			anchor := col.DOM.Find("form > a")
			if anchor.Length() > 0 {
				data = anchor.Contents().First().Text()
			} else {
				data = col.DOM.Contents().First().Text()
			}
			rowData = append(rowData, data)

			count++
		})

		var course model.Course
		timetableID := rowData[0]
		if timetableID != "-" {
			fmt.Sscan(timetableID, course.TimetableID)
		}

		title := rowData[1]
		fmt.Sscan(title, &course.Title)

		class := rowData[2]
		if class != "-" {
			fmt.Sscan(class, course.Class)
		}

		t := rowData[3]
		fmt.Sscan(t, &course.Type)

		credits := rowData[4]
		fmt.Sscan(credits, &course.Credits)

		result = append(result, course)
	})

	return result
}
