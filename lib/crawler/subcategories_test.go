package crawler

import (
	"amazon/lib/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)



func TestSubCategories(t *testing.T) {
	var passed bool
	var numberofPassedCategories int
	var departments []types.Department
	data, err := ioutil.ReadFile("/home/ubuntu/Documents/amazon/testData/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)
	departments = GetSubCategories(departments)
	for _, department := range departments {
		for _, category := range department.Categories {
			if len(category.SubCategories) >= 1 {
				passed = true
				numberofPassedCategories++
			}
		}
	}
	if !passed {
		t.Error("Test Failed")
	}else {
		data := JsonMarshal(departments)
		err := ioutil.WriteFile("/home/ubuntu/Documents/amazon/testData/deps.json", data, 0755)
		HandleError(err)
		fmt.Println(numberofPassedCategories)
	}
}