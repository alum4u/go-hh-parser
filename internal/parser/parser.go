package parser

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"fmt"
	"regexp"
)



func Parse(text string) {

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(text))

	
	doc.Find(".vacancy-serp-item-body").Each( func(i int, el *goquery.Selection) {
		cost := el.Find("[data-qa=vacancy-serp__vacancy-compensation]").Text()
		processCost(cost)
	});

}

func processCost(cost string) string {
	if !strings.Contains(cost, "руб") {
		return ""
	}
	formatted := regexp.MustCompile(`[\d]{1,3}`)

	partedCost := formatted.FindAllString(cost, 2)
	
	return strings.Join(partedCost, "")	
}