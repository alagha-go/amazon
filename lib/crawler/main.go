package crawler

import (
	department"amazon/lib/departments"
	"bytes"
	"log"
	"encoding/json"
	"time"
)


func Main() {
	for {
		GetAllDepartments()
		time.Sleep(72*time.Hour)
	}
}


func GetAllDepartments() {
	departments := GetDepartments()

	for dindex, department := range departments {
		departments[dindex] = GetSubCategories(department)
		for cindex, category := range departments[dindex].Categories {
			for sindex, SubCategory := range category.SubCategories {
				departments[dindex].Categories[cindex].SubCategories[sindex] = GetTypes(SubCategory)
				SubCategory = departments[dindex].Categories[cindex].SubCategories[sindex]
				for tindex, Type := range SubCategory.Types {
					departments[dindex].Categories[cindex].SubCategories[sindex].Types[tindex] = GetSubTypes(Type)
				}
			}
		}
	}

	err := department.InsertToDB(departments)
	if err != nil {
		GetAllDepartments()
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