package main 

import (
	"fmt"
	"github.com/gocolly/colly"
	"encoding/csv"
  "log"
  "os"
)

type Industry struct {
	Url, Image, Name string
}

var industries []Industry

func main() {
	fmt.Println("Hello World")

	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	c.Visit("https://brightdata.com/")

	c.OnHTML(".elementor-element-6b05593c .section_cases__item", func(e *colly.HTMLElement) {
    url := e.Attr("href")
    image := e.ChildAttr(".elementor-image-box-img img", "data-lazy-src")
    name := e.ChildText(".elementor-image-box-content .elementor-image-box-title")
    if url!= ""  || image != ""  || name != ""  {
			industry := Industry{
				Url:   url,
				Image: image,
				Name:  name,
			}

      industries = append(industries, industry)
    }
  })

	// open the output CSV file
	file, err := os.Create("industries.csv")
	// if the file creation fails
	if err != nil {
		log.Fatalln("Failed to create the output CSV file", err)
	}
	// release the resource allocated to handle
	// the file before ending the execution
	defer file.Close()

  // create a CSV file writer
   writer := csv.NewWriter(file)
	// release the resources associated with the
	// file writer before ending the execution
	defer writer.Flush()

	// add the header row to the CSV
	headers := []string{
		"url",
		"image",
		"name",
	}
	writer.Write(headers)

	// store each Industry product in the
	// output CSV file
	for _, industry := range industries {
		// convert the Industry instance to
		// a slice of strings
		record := []string{
				industry.Url,
				industry.Image,
				industry.Name,
		}

		// add a new CSV record
		writer.Write(record)
	}
}