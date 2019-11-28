package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mmcdole/gofeed"
)

func main() {
	// Connect with postgreSQL Database
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=biorxiv user=postgres password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	// Get research papers info from RSS
	getRSS(db)

	// Close the connection with DB
	defer db.Close()
}

type Articles struct {
	Title       string
	Author      string
	Link        string
	Description string
	Published   string
	Doi         string
}

// type BioRxivData []*BioRxivArticle

func getRSS(db *gorm.DB) {
	url := "http://connect.biorxiv.org/biorxiv_xml.php?subject=genomics+bioinformatics"

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	items := feed.Items

	// var allData BioRxivData

	for _, item := range items {
		data := new(Articles)
		data.Title = strings.Replace(item.Title, "\n", "", -1)
		data.Link = item.Link // BioRxiv article URL
		data.Description = strings.Replace(item.Description, "\n", "", -1)
		data.Author = item.Author.Name
		data.Published = item.Published
		extension := item.Extensions
		doi := extension["dc"]["identifier"]
		data.Doi = strings.Replace(doi[0].Value, "doi:", "", -1) // DOI

		// Insert article info into database if the article does not exist
		if err := db.Create(&data); err.Error != nil {
			fmt.Println(err.Error)
		}
		// allData = append(allData, data)
	}

	// fmt.Println(allData)
}
