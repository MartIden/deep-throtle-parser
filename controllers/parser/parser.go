package parser

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func GetInnerLinks(template *goquery.Document, baseLink string) []string {

	var results []string
	template.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		isMatch, _ := regexp.MatchString(baseLink, href)
		if isMatch {
			results = append(results, href)
		}
	})
	return results
}
