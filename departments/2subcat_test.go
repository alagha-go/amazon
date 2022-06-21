package departments

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)


func TestSubCategories(t *testing.T) {
	var passed bool
	var departments []Department
	data, err := ioutil.ReadFile("../DB/deps.json")
	HandleError(err)
	err = json.Unmarshal(data, &departments)
	HandleError(err)
	for index := range departments {
		if departments[index].Title == "electronics" {
			departments[index].SetCategories()
			for i := range departments[index].Categories {
				if len(departments[index].Categories[i].SubCategories) > 0 {
					passed = true
					break
				}
			}
			break
		}
	}
	
	if !passed {
		t.Error("failed to collect subcategories")
	}
	data = JsonMarshal(departments)
	ioutil.WriteFile("../DB/deps.json", data, 0755)
}