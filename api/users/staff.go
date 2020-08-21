package users

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Employee struct {
	profilePic  string
	fullName    string
	mailAdress  string
	description string
}

func GetEmployees() []Employee {
	var employees []Employee
	c := colly.NewCollector(colly.MaxDepth(1))

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Requesting on", request.URL)
	})

	c.OnHTML("h2", func(element *colly.HTMLElement) {
		employees = append(employees, Employee{fullName: element.Text})
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
