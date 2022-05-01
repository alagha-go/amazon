package departments

import (
	"amazon/lib/types"
	"context"
)


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