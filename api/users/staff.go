package users

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

type Employee struct {
	ProfilePic  string `json:"profile_pic"`
	FullName    string `json:"full_name"`
	MailAdress  string `json:"mail_adress"`
	Description string `json:"description"`
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
		var profileLink = element.ChildAttr("img", "src")
		var mailndescription = strings.Split(element.ChildText(".sitotheque p"), "\n")

		if len(profileLink) != 0 && !strings.Contains(profileLink, "twitter") {
			employees = append(employees, Employee{
				FullName:    element.ChildText("h2"),
				MailAdress:  mailndescription[0],
				ProfilePic:  "https://camelia55.meuse.fr" + element.ChildAttr("img", "src"),
				Description: mailndescription[1],
			})
		}
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
