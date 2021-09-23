package parser

import (
	"regexp"
	"github.com/PuerkitoBio/goquery"
)


func GetInnerLinks(template *goquery.Document, baseLink string,c chan []string) {

	var results []string
	template.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		isMatch, _ := regexp.MatchString(baseLink, href)
		if isMatch {
			results = append(results, href)
		}
	})
	c <- results
}