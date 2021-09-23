package parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
)

type Link string

func GetInnerLinks(template io.ReadCloser, baseLink string) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(template)
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
	defer template.Close()
	return results, nil
}
