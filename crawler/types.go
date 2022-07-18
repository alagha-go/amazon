package crawler

import (
	// "time"

	"fmt"
	"strings"

	"github.com/gocolly/colly"
)


func (subcategory *SubCategory) SetTypes(){
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
					var Type Type
					Type.Url = "https://amazon.com" + element.Attr("href")
					Type.Title = element.Text
					Type.Title = strings.ReplaceAll(Type.Title, "\n", "")
					Type.Title = strings.ReplaceAll(Type.Title, "  ", " ")
					Type.Title = strings.TrimSpace(Type.Title)
					subcategory.Types = append(subcategory.Types, Type)
					return false
				}
				return true
			})
		}
		})
	})
	
	collector.Visit(subcategory.Url)
	Reload()
}