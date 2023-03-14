package parser

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
)



func Parse(text string) {

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(text))

	
	doc.Find(".vacancy-serp-item-body").Each( func(i int, el *goquery.Selection) {
		fmt.Println(el.Text())
	});

}