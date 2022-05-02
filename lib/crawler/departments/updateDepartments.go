package departments

import (
	"amazon/lib/departments"
	"amazon/lib/types"
	"encoding/json"
)


func UpdateDepartments() {
	var Departments []types.Department
	data, _ := departments.GetDepartments()
	json.Unmarshal(data, &Departments)

	if len(Departments) < 1 {
		UpdateDepartments()
	}

	for _, Department := range Departments {
		for cindex, Category := range Department.Categories {
			if len(Category.SubCategories) >= 1 {
				continue
			}
			Department.Categories[cindex].SubCategories = CollectSubCategories(Category)
			Category.SubCategories = Department.Categories[cindex].SubCategories
			for sindex, SubCategory := range Category.SubCategories {
				if len(SubCategory.Types) >= 1 {
					continue
				}
				Department.Categories[cindex].SubCategories[sindex] = GetTypes(SubCategory)
				SubCategory = Department.Categories[cindex].SubCategories[sindex]
				for tindex, Type := range SubCategory.Types {
					if len(Type.SubTypes) >= 1 {
						continue
					}
					Department.Categories[cindex].SubCategories[sindex].Types[tindex] = GetSubTypes(Type)
					Type = Department.Categories[cindex].SubCategories[sindex].Types[tindex]
				}
			}
		}

		departments.UpdateDepartment(Department)
	}
}