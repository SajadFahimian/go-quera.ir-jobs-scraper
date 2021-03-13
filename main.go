package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Job a data struct for select data
type Job struct {
	Title   string `json:"title"`
	Company string `json:"company"`
	Collab  string `json:"collab"`
	Salary  string `json:"salary"`
	Remote  string `json:"remote"`
}

// Numbe page
var Page = 0

// This func for get number of page
func NumberOfPages() {
	collector := colly.NewCollector(colly.AllowedDomains("quera.ir", "www.quera.ir"))
	collector.OnHTML(".center-text .item", func(e *colly.HTMLElement) {

		if e.Name == "a" {
			value, err := strconv.Atoi(strings.Split(e.Attr("href"), "=")[1])
			if err != nil {
				return
			} // End if
			if value > Page {
				Page = value
			} // End if

		} // End if
	}, // End func
	)
	collector.Visit("https://www.quera.ir/careers/jobs")
}

// Main func for read all page and save data from jobs array
func main() {
	NumberOfPages()
	jobs := make([]Job, 0)
	collector := colly.NewCollector(colly.AllowedDomains("quera.ir", "www.quera.ir"))
	for i := 1; i < Page; i++ {
		collector.OnHTML("#jobs-segment .content", func(e *colly.HTMLElement) {

			job := Job{
				Title:   strings.TrimSpace(e.ChildText("h2 span")),
				Company: strings.TrimSpace(e.ChildText(".meta")),
				Collab:  strings.TrimSpace(e.ChildText(".extra .job-collab")),
				Salary:  strings.TrimSpace(e.ChildText(".extra .job-salary")),
				Remote:  strings.TrimSpace(e.ChildText(".extra .job-remote")),
			} // End data struct
			jobs = append(jobs, job)

		}, // End func
		)
		collector.Visit("https://www.quera.ir/careers/jobs")
		collector.Visit(fmt.Sprintf("https://www.quera.ir/careers/jobs?page=%d", i))
		fmt.Printf("Page %d ...\n", i)
	} // End for loop
	fmt.Println("Reade All page")
	writedJson(jobs)

}

// This func for write all data from now.json
func writedJson(data []Job) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile(fmt.Sprintf("data(%v).json", time.Now().UTC().UnixNano()), file, 0644)
	fmt.Println("Successfully create json data")
}
