package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
}

func main() {
	// file, err := os.OpenFile("result.html", os.O_RDONLY|os.O_CREATE, 0666)

	doc, err := goquery.NewDocument("https://www.google.com/search?q=%E4%B8%AD%E9%87%8E%E4%BA%8C%E4%B9%83&source=lnms&tbm=isch&sa=X&ved=2ahUKEwjYy6_egK_mAhXaMt4KHd0NBe4Q_AUoAXoECAoQAw&biw=412&bih=766")
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	res, err := doc.Find("body").Html()
	if err != nil {
		fmt.Print("dom get failed")
	}
	ioutil.WriteFile("./result.html", []byte(res), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, _ := ioutil.ReadFile("./result.html")
	stringReader := strings.NewReader(string(fileInfos))
	doc, err = goquery.NewDocumentFromReader(stringReader)
	if err != nil {
		fmt.Print("url scarapping failed")
	}
	// // Load the HTML document

	doc.Find("img").Each(func(index int, s *goquery.Selection) {
		fmt.Println(s.Attr("src"))
	})
}
