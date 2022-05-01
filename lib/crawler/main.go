package crawler

import (
	"amazon/lib/types"
	"bytes"
	"log"
	"context"
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

	err := InsertToDB(departments)
	if err != nil {
		GetAllDepartments()
	}
}


func InsertToDB(departments []types.Department) error {
	var documents []interface{}
	ctx := context.Background()
	collection := types.Client.Database("Amazon").Collection("Departments")

	for _, department := range departments {
		documents = append(documents, department)
	}

	_, err := collection.InsertMany(ctx, documents)
	return err
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