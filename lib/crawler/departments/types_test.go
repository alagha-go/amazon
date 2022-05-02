package departments

import (
	"amazon/lib/types"
	"encoding/json"
	"io/ioutil"
	"testing"
)


func TestGetTypes(t *testing.T) {
	var passed bool
	var departments []types.Department
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)

	for dindex, department := range departments {
		for cindex, category := range department.Categories {
			for sindex, subcategory := range category.SubCategories {
				departments[dindex].Categories[cindex].SubCategories[sindex] = GetTypes(subcategory)
				if len(departments[dindex].Categories[cindex].SubCategories[sindex].Types) > 0 {
					passed = true
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