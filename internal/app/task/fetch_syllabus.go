package task

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
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
	courses := model.Courses{}
	baseURL := fmt.Sprintf("%s://%s/", w.targetURL.Scheme, w.targetURL.Host)

	// colly, goquery用いてスクレイピング
	c := colly.NewCollector(
		colly.AllowedDomains(w.targetURL.Host),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{Parallelism: 10})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML("#search_result table.data_list_tbl", func(e *colly.HTMLElement) {
		res := scrapeCourse(e)
		courses = append(courses, res...)
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

	// for i, c := range courses {
	// 	fmt.Printf("%d件目\n", i)
	// 	if c.TimetableID == nil {
	// 		fmt.Printf("\t%v\n", nil)
	// 	} else {
	// 		fmt.Printf("\t%v\n", *c.TimetableID)
	// 	}
	// 	fmt.Printf("\t%v\n", c.Title)
	// 	if c.Class == nil {
	// 		fmt.Printf("\t%v\n", nil)
	// 	} else {
	// 		fmt.Printf("\t%v\n", *c.Class)
	// 	}
	// 	fmt.Printf("\t%v\n", c.Type)
	// 	fmt.Printf("\t%v\n", c.Credits)
	// 	fmt.Printf("\t%v\n", c.Instructors)
	// 	fmt.Printf("\t%v\n", c.Year)
	// 	fmt.Printf("\t%v\n", c.Semester)
	// 	if c.Day == nil {
	// 		fmt.Printf("\t%v\n", nil)
	// 	} else {
	// 		fmt.Printf("\t%v\n", *c.Day)
	// 	}
	// }

	n, err := w.repository.BulkInsertCourses(courses)
	if err != nil {
		return err
	}

	log.Printf("bulk inserted %d courses\n", n)

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
			course.TimetableID = &timetableID
		}

		title := rowData[1]
		course.Title = title

		class := rowData[2]
		if class != "-" {
			course.Class = &class
		}

		t := rowData[3]
		course.Type = t

		credits := rowData[4]
		course.Credits, _ = strconv.Atoi(credits)

		instructors := rowData[5]
		course.Instructors = instructors

		year := rowData[6]
		course.Year = year

		semester := rowData[7]
		course.Semester = semester

		day := rowData[8]
		if day != "-" {
			course.Day = &day
		}

		result = append(result, course)
	})

	return result
}
