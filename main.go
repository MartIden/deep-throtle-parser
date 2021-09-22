package main

import (
	"fmt"
	"github.com/MartIden/deep-throtle-parser/command"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"regexp"
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
			resp, err := getPage(link)
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			links, err := getInnerLinks(resp, args.Url)
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

func getInnerLinks(resp *http.Response, baseLink string) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	var results []string
	if err != nil {
		return nil, err
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {

		href, _ := s.Attr("href")
		isMatch, _ := regexp.MatchString(baseLink, href)
		if isMatch {
			results = append(results, href)
		}
	})
	defer resp.Body.Close()
	return results, nil
}

func getPage(uri string) (*http.Response, error) {
	client := http.Client{}
	resp, err := client.Get(uri)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
