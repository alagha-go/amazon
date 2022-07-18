package crawler

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)


func TestGetSubTypes(t *testing.T) {
	println("called Test3")
	var passed bool
	var departments []Department
	data, err := ioutil.ReadFile("../DB/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)

	loop:
	for dindex, department := range departments {
		if department.Title == "electronics" {
			for cindex, category := range department.Categories {
				for sindex, subcategory := range category.SubCategories {
					departments[dindex].Categories[cindex].SubCategories[sindex].SetTypes()
					subcategory = departments[dindex].Categories[cindex].SubCategories[sindex]
					for tindex := range subcategory.Types {
						departments[dindex].Categories[cindex].SubCategories[sindex].Types[tindex].SetSubTypes()
						if len(subcategory.Types[tindex].SubTypes) > 0 {
							passed = true
							break loop
						}
					}
				}
			}
		}
	}

	if !passed {
		t.Error("failed to collect types")
		data = JsonMarshal(departments)
		ioutil.WriteFile("../DB/deps.json", data, 0755)
	}
}