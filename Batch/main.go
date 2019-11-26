package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {

	url := "http://connect.biorxiv.org/biorxiv_xml.php?subject=genomics+bioinformatics"

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)

	items := feed.Items

	for _, item := range items {
		fmt.Println(item.Title)
		fmt.Println(item.Link) // BioRxiv article URL
		fmt.Println(item.Description)
		fmt.Println(item.Author)
		fmt.Println(item.Published)

		// DOI
		extension := item.Extensions
		doi := extension["dc"]["identifier"]
		fmt.Println(doi[0].Value)
		break
	}
}
