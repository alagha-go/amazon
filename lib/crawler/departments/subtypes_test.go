package departments

import (
	"amazon/lib/types"
	"encoding/json"
	"io/ioutil"
	"testing"
)


func TestGetSubTypes(t *testing.T) {
	var passed bool
	var departments []types.Department
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)

	for dindex, department := range departments {
		for cindex, category := range department.Categories {
			for sindex, subcategory := range category.SubCategories {
				for tindex, Type := range subcategory.Types {
					departments[dindex].Categories[cindex].SubCategories[sindex].Types[tindex] = GetSubTypes(Type)
					if len(departments[dindex].Categories[cindex].SubCategories[sindex].Types[tindex].SubTypes) > 0 {
						passed = true
					}
				}
			}
		}
	}

	if !passed {
		t.Error("failed to collect subcategories")
	}
	data = JsonMarshal(departments)
	ioutil.WriteFile("/home/ubuntu/Documents/amazon/testData/deps.json", data, 0755)
}