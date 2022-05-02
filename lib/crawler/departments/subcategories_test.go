package departments

import (
	"amazon/lib/types"
	"encoding/json"
	"io/ioutil"
	"testing"
)


func TestSubCategories(t *testing.T) {
	var passed bool
	var departments []types.Department
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)
	for index, department := range departments {
		department = GetSubCategories(department)
		if len(department.Categories[0].SubCategories) > 0 {
			passed = true
		}
		departments[index] = department
	}
	
	if !passed {
		t.Error("failed to collect subcategories")
	}
	data = JsonMarshal(departments)
	ioutil.WriteFile("/home/ubuntu/Documents/amazon/testData/deps.json", data, 0755)
}