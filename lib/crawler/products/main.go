package products

import (
	"amazon/lib/departments"
	"amazon/lib/types"
	"encoding/json"
)


func Main() {
	var Departments []types.Department
	data, _ := departments.GetDepartments()
	json.Unmarshal(data, &Departments)
	CollcetProducts(Departments)
}