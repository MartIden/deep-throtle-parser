package main

import (
	"log"
	"os"

	"github.com/MartIden/deep-throtle-parser/command"
	"github.com/MartIden/deep-throtle-parser/parser"
	"github.com/MartIden/deep-throtle-parser/request"
)

func main() {

	args, err := command.ProcessArgs(os.Args)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	var linksForParsing [][]string
	linksForParsing = append(linksForParsing, []string{args.Url})

	for _, links := range linksForParsing {
		var currentLinks []string
		
		for _, link := range links {
			resp, err := request.GetPageByUri(link)

			template := resp.Body
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			links, err := parser.GetInnerLinks(template, args.Url)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			for _, currentLink := range links {
				currentLinks = append(currentLinks, currentLink)
			}
		}
		linksForParsing = append(linksForParsing, currentLinks)
	}

	var listOfLinks []string

	for _, resultLinks := range linksForParsing {
		for _, link := range resultLinks {
			listOfLinks = append(listOfLinks, link)
		}
	}
}
