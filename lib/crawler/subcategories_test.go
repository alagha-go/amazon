package crawler

import (
	"amazon/lib/types"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"testing"
)


func TestSubCategories(t *testing.T) {
	var passed bool
	var passedNo int
	var departments []types.Department
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)
	for index, department := range departments {
		department = GetSubCategories(department)
		if len(department.Categories[0].SubCategories) > 0 {
			passed = true
			passedNo++
		}
		departments[index] = department
	}
	data = JsonMarshal(departments)
	ioutil.WriteFile("/home/ubuntu/Documents/amazon/testData/deps.json", data, 0755)

	if !passed {
		t.Error("failed to collect subcategories")
	}else{
		fmt.Println(passedNo)
	}
}