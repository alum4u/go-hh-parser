package parser

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"regexp"
)



func Parse(text string) []string {

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(text))

	costs := make([]string,0)

	doc.Find(".vacancy-serp-item-body").Each( func(i int, el *goquery.Selection) {
		cost := el.Find("[data-qa=vacancy-serp__vacancy-compensation]").Text()

		costs = append(costs,processCost(cost))
	});
	
	return costs
}

func processCost(cost string) string {
	if !strings.Contains(cost, "руб") {
		return ""
	}
	formatted := regexp.MustCompile(`[\d]{1,3}`)

	partedCost := formatted.FindAllString(cost, 2)
	
	return strings.Join(partedCost, "")	
}