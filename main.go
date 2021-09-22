package main

import (
	"fmt"
	"github.com/MartIden/deep-throtle-parser/command"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

func main() {

	_, err := command.ProcessArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	baseLink := "https://wpnew.ru/blog"
	resp, _ := getPage(baseLink)
	getInnerLinks(resp)
}

func getInnerLinks(resp *http.Response) {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)

	})
}

func getPage(uri string) (*http.Response, error) {
	client := http.Client{}
	resp, err := client.Get(uri)
	if err != nil {
		fmt.Println(err)
		return resp, err
	}
	defer resp.Body.Close()
	return resp, nil
}
