package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {

	blogTitles, err := GetLatestBlogTitles("https://golangcode.com")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Blog Titles:")
	fmt.Println(blogTitles)
}

func GetLatestBlogTitles(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	titles := ""
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		titles += "-" + s.Text() + "\n"
	})
	return titles, nil
}
