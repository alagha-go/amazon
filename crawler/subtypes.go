package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)


func (Type *Type) SetSubTypes() {
	collector := colly.NewCollector()

	err := collector.SetProxy(ProxyUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	collector.OnHTML("body", func(element *colly.HTMLElement) {
		element.ForEach("li", func(_ int, element *colly.HTMLElement) {
			if element.Attr("class") == "a-spacing-micro s-navigation-indent-2" {
			element.ForEachWithBreak("a", func(_ int, element *colly.HTMLElement) bool {
				off := element.Attr("data-routing")
				if off == "off" {
					var subType SubType
					subType.Url = "https://amazon.com" + element.Attr("href")
					subType.Title = element.Text
					subType.Title = strings.ReplaceAll(subType.Title, "\n", "")
					subType.Title = strings.ReplaceAll(subType.Title, "  ", " ")
					subType.Title = strings.TrimSpace(subType.Title)
					Type.SubTypes = append(Type.SubTypes, subType)
					return false
				}
				return true
			})
		}
		})
	})
	
	collector.Visit(Type.Url)
	Reload()
}