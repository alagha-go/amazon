package crawler

import (
	departmentspkg "amazon/lib/departments"
	"amazon/lib/types"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"time"
)


func Main() {
	for {
		CollectAllDepartments()
		for index:=0; index<5; index++ {
			UpdateDepartments()
		} 
		time.Sleep(720*time.Hour)
	}
}


func CollectAllDepartments() {
	collection := types.Client.Database("Amazon").Collection("Departments")
	collection.Drop(context.Background())
	departments := GetDepartments()

	for _, department := range departments {
		department = GetSubCategories(department)
		for cindex, category := range department.Categories {
			for sindex, SubCategory := range category.SubCategories {
				department.Categories[cindex].SubCategories[sindex] = GetTypes(SubCategory)
				SubCategory = department.Categories[cindex].SubCategories[sindex]
				for tindex, Type := range SubCategory.Types {
					department.Categories[cindex].SubCategories[sindex].Types[tindex] = GetSubTypes(Type)
				}
			}
		}
		departmentspkg.InsertOneDepartment(department)
	}
}


func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}


func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}