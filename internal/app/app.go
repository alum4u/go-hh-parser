package app

import (
	"fmt"
	"main/internal/grabber"
	"main/internal/parser"
)

func Start() {
	fmt.Println("App started")

	response, _ := grabber.Grab("https://hh.ru/search/vacancy?text=golang&area=1")

	
	fmt.Println(parser.Parse(response))
	
}