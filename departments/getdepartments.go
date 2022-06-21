package departments

import (
	"strings"

	"github.com/gocolly/colly"
)

const (
	DepartmentsPageUrl = "https://www.amazon.com/gp/navigation/ajax/generic.html?ajaxTemplate=hamburgerMainContent&pageType=Gateway&hmDataAjaxHint=1&navDeviceType=desktop&isSmile=0&isPrime=0&isBackup=false&hashCustomerAndSessionId=ffbb1810709292f57a2507d0669c66868eea6018&isExportMode=true&languageCode=en_US&environmentVFI=AmazonNavigationCards%2Fdevelopment%40B6079054483-AL2_x86_64&secondLayerTreeName=prm_digital_music_hawkfire%2Bkindle%2Bandroid_appstore%2Belectronics_exports%2Bcomputers_exports%2Bsbd_alexa_smart_home%2Barts_and_crafts_exports%2Bautomotive_exports%2Bbaby_exports%2Bbeauty_and_personal_care_exports%2Bwomens_fashion_exports%2Bmens_fashion_exports%2Bgirls_fashion_exports%2Bboys_fashion_exports%2Bhealth_and_household_exports%2Bhome_and_kitchen_exports%2Bindustrial_and_scientific_exports%2Bluggage_exports%2Bmovies_and_television_exports%2Bpet_supplies_exports%2Bsoftware_exports%2Bsports_and_outdoors_exports%2Btools_home_improvement_exports%2Btoys_games_exports%2Bvideo_games_exports%2Bgiftcards%2Bamazon_live%2BAmazon_Global"
)


func GetDepartments() []Department {
	collector := colly.NewCollector()
	var Departments []Department

	collector.OnHTML("body", func(element *colly.HTMLElement) {
		element.ForEach(".hmenu.hmenu-translateX-right", func(index int, element *colly.HTMLElement) {
			Departments = append(Departments, GetDepartment(element))
		})
	})

	collector.Visit(DepartmentsPageUrl)
	return Departments
}

func GetDepartment(element *colly.HTMLElement) Department {
	var Department Department
	element.ForEach(".hmenu-item", func(index int, element *colly.HTMLElement){
		if index >= 2 {
			var category Category
			category.Title = element.Text
			url := element.Attr("href")
			if url != "" {
				if !strings.Contains(url, "https://"){
					category.Url = "https://amazon.com" + url
				}else {
					category.Url = url
				}
				Department.Categories = append(Department.Categories, category)
			}
		}else if index == 1 {
			Department.Title = element.Text
		}
	})

	return Department
}