package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MartIden/deep-throtle-parser/pkg/command"
	"github.com/MartIden/deep-throtle-parser/pkg/parser"
	"github.com/MartIden/deep-throtle-parser/pkg/request"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	args, err := command.ProcessArgs(os.Args)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	var linksForParsing [][]string
	linksForParsing = append(linksForParsing, []string{args.Url})

	for i := 0; i < int(args.Deep); i++ {
		var currentLinks []string
		for _, link := range linksForParsing[i] {
			resp, err := request.GetPageByUri(link)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			pageLinks := parser.GetInnerLinks(doc, args.Url)
			fmt.Println(pageLinks)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			for _, currentLink := range pageLinks {
				currentLinks = append(currentLinks, currentLink)
			}
		}
		linksForParsing = append(linksForParsing, currentLinks)
	}

	fmt.Println(linksForParsing)
}
