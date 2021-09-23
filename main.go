package main

import (
	"fmt"
	"github.com/MartIden/deep-throtle-parser/controllers/command"
	"github.com/MartIden/deep-throtle-parser/controllers/parser"
	"github.com/MartIden/deep-throtle-parser/controllers/request"
	"log"
	"os"
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
			for _, item := range links {
				fmt.Println(item)
			}
		}

	}
}
