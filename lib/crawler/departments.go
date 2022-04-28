package crawler

import (
	"amazon/lib/types"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DepartmentsPageUrl = "https://www.amazon.com/gp/navigation/ajax/generic.html?ajaxTemplate=hamburgerMainContent&pageType=Gateway&hmDataAjaxHint=1&navDeviceType=desktop&isSmile=0&isPrime=0&isBackup=false&hashCustomerAndSessionId=ffbb1810709292f57a2507d0669c66868eea6018&isExportMode=true&languageCode=en_US&environmentVFI=AmazonNavigationCards%2Fdevelopment%40B6079054483-AL2_x86_64&secondLayerTreeName=prm_digital_music_hawkfire%2Bkindle%2Bandroid_appstore%2Belectronics_exports%2Bcomputers_exports%2Bsbd_alexa_smart_home%2Barts_and_crafts_exports%2Bautomotive_exports%2Bbaby_exports%2Bbeauty_and_personal_care_exports%2Bwomens_fashion_exports%2Bmens_fashion_exports%2Bgirls_fashion_exports%2Bboys_fashion_exports%2Bhealth_and_household_exports%2Bhome_and_kitchen_exports%2Bindustrial_and_scientific_exports%2Bluggage_exports%2Bmovies_and_television_exports%2Bpet_supplies_exports%2Bsoftware_exports%2Bsports_and_outdoors_exports%2Btools_home_improvement_exports%2Btoys_games_exports%2Bvideo_games_exports%2Bgiftcards%2Bamazon_live%2BAmazon_Global"
)


func GetDepartments() []types.Department {
	collector := colly.NewCollector()
	var Departments []types.Department

	collector.OnHTML("body", func(element *colly.HTMLElement) {
		element.ForEach(".hmenu.hmenu-translateX-right", func(index int, element *colly.HTMLElement) {
			var department types.Department
			department.ID = primitive.NewObjectID()
			department.Title = element.ChildText(".hmenu-item.hmenu-title")
			department.Categories = GetCategories(element)
			Departments = append(Departments, department)
		})
	})

	collector.Visit(DepartmentsPageUrl)
	return Departments
}

func GetCategories(element *colly.HTMLElement) []types.Category {
	var categories []types.Category
	element.ForEach(".hmenu-item", func(index int, element *colly.HTMLElement){
		var category types.Category
		category.ID = primitive.NewObjectID()
		category.Title = element.Text
		category.Url = "https://amazon.com" + element.Attr("href")
		categories = append(categories, category)
	})

	return categories
}