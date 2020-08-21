package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

type Employee struct {
	profilePic  string
	fullName    string
	mailAdress  string
	description string
}

func GetEmployees() []Employee {
	var employees []Employee
	c := colly.NewCollector(
		colly.MaxDepth(2))

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 3 * time.Second,
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Requesting on", request.URL)
	})

	c.OnHTML(".article_content", func(element *colly.HTMLElement) {
		var link string
		if strings.Contains(element.ChildAttr("a", "href"), "mailto") {
			link := element.ChildAttr("a", "href")
			link = strings.ReplaceAll(link, "mailto:", "")
		}

		employees = append(employees, Employee{
			fullName:    element.ChildText("h2"),
			mailAdress:  link,
			profilePic:  "https://camelia55.meuse.fr" + element.ChildAttr("img", "src"),
			description: element.ChildText(".sitotheque p"),
		})
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("Error on requested url", response.Request.URL, "with code", response, "\n(", err, ")")
	})

	err := c.Visit("https://camelia55.meuse.fr/equipe")
	if err != nil {
		panic(err)
	}
	return employees
}
